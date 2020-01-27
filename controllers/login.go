package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/denerFernandes/goreststore/auth"
	"github.com/denerFernandes/goreststore/models"
	"github.com/denerFernandes/goreststore/responses"
	"github.com/denerFernandes/goreststore/utils"
)

// Login / Make login
func (controller *Controller) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	err = models.VerifyPassword(user.Password, user.Password)
	if err != nil {
		formattedError := utils.Error(err.Error())
		responses.ERROR(w, r, http.StatusUnprocessableEntity, formattedError)
		return
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		responses.ERROR(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, r, http.StatusOK, "Login Successfully", token)
}
