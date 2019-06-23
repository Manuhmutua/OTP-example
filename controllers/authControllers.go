package controllers

import (
	"encoding/json"
	"github.com/Manuhmutua/movies-backend-apis/models"
	u "github.com/Manuhmutua/movies-backend-apis/utils"
	"github.com/xlzd/gotp"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	totp := gotp.NewDefaultTOTP(gotp.RandomSecret(16))
	totp.ProvisioningUri("OurMesseger", "movieShow")

	resp := models.Login(account.Phone, account.OTP, account.Verified, totp)
	u.Respond(w, resp)
}
