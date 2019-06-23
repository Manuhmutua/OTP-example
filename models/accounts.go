package models

import (
	"encoding/json"
	"fmt"
	u "github.com/Manuhmutua/movies-backend-apis/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/xlzd/gotp"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uuid.UUID
	Phone  string
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"primary_key;auto_increment:false"`
	Phone    string    `json:"phone_number"`
	UserName string    `json:"user_name"`
	OTP      string    `json:"otp_number"`
	Token    string    `json:"token";sql:"-"`
	Active   bool      `json:"active"`
}

var totp *gotp.TOTP

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Phone, "+") {
		return u.Message(false, "Phone Number address is required"), false
	}

	if len(account.UserName) < 3 {
		return u.Message(false, "Username is required"), false
	}

	//PhoneNumber must be unique
	temp := &Account{}

	//check for errors and duplicate phones
	err := GetDB().Table("accounts").Where("phone = ?", account.Phone).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"+err.Error()), false
	}
	if temp.Phone != "" {
		return u.Message(false, "Phone Number address already in use by another user."), false
	}

	//check for errors and duplicate username
	err = GetDB().Table("accounts").Where("user_name = ?", account.UserName).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"+err.Error()), false
	}
	if temp.UserName != "" {
		response := fmt.Sprintf("Username: %d is already in use by another user.", account.UserName)
		return u.Message(false, response), false
	}

	return u.Message(false, "Requirement passed"), true
}

func sendMessage(userName string, phoneNumber string, otp *gotp.TOTP) map[string]interface{} {
	accountSid := os.Getenv("SMS_ACCOUNT_SID")
	authToken := os.Getenv("SMS_AUTH_TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Set up rand
	rand.Seed(time.Now().Unix())

	msgData := url.Values{}
	msgData.Set("To", phoneNumber)
	msgData.Set("From", os.Getenv("SMS_ACCOUNT_NUMBER"))
	msgData.Set("Body", "Hello, "+userName+" . Your OTP pin is: "+otp.Now())
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		return u.Message(false, "Failed to create account, connection error.(Sending Message)")
	}
	return nil
}

func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	// Generate UUID
	Uuid, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		return u.Message(false, "Failed to create account, connection error.(UUID)")
	}
	account.UUID = Uuid

	// Handle Message Logic
	totp := gotp.NewDefaultTOTP(gotp.RandomSecret(16))
	totp.ProvisioningUri("OurMesseger", "movieShow")

	sendMessage(account.UserName, account.Phone, totp)

	GetDB().Create(account)

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.UUID, Phone: account.Phone}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
	account.Token = tokenString
	account.Active = false

	response := u.Message(true, "Account has been created, Proceed With Verification")
	response["account"] = account
	return response
}

func (account *Account) Login(phone, otp string, status bool) map[string]interface{} {

	err := GetDB().Table("accounts").Where("phone = ?", phone).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Phone Number address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	if otp != totp.Now() { //OTP does not match!
		return u.Message(false, "Invalid otp. Please try again")
	}

	//Create JWT token
	tk := &Token{UserId: account.UUID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
	account.Token = tokenString //Store the token in the response
	account.Active = true

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func (account *Account) GetUser(u uuid.UUID) *Account {
	GetDB().Table("accounts").Where("uuid = ?", u).First(account)
	if account.Phone == "" { //User not found!
		return nil
	}
	return account
}
