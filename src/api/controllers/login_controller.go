package controllers

import (
	"back-account/src/api/auth"
	"back-account/src/api/models"
	"back-account/src/api/responses"
	"encoding/json"
	"errors"
	"net/http"
)

// Login is the signIn method
func Login(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("packet is recived to server .."))
		return
	}

	// err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, user, err := auth.SignIn(user.UserName, user.Password)

	if err != nil {
		responses.ERROR(w, http.StatusNetworkAuthenticationRequired, errors.New("نام کاربری یا رمز عبور اشتباه می باشد"))
		return

	}

	newtoken := res{Yourtoken: token, YourUser: user}

	responses.JSON(w, http.StatusOK, &newtoken)
}

type res struct {
	Yourtoken string      `json:"token"`
	YourUser  models.User `json:"user"`
}
