package controllers

import (
	"encoding/json"
	"github.com/Manuhmutua/movies-backend-apis/models"
	u "github.com/Manuhmutua/movies-backend-apis/utils"
	"github.com/xlzd/gotp"
	"net/http"
)

func getTotp() *gotp.TOTP {
	totp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	totp.ProvisioningUri("OurMesseger", "movieShow")
	return totp
}

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create(getTotp()) //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Phone, account.OTP, getTotp())
	u.Respond(w, resp)
}

var Reset = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Reset(account.Phone, getTotp())
	u.Respond(w, resp)
}