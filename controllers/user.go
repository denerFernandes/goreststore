package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/denerFernandes/goreststore/database"
	"github.com/denerFernandes/goreststore/models"
	"github.com/denerFernandes/goreststore/responses"
	"github.com/denerFernandes/goreststore/utils"
	"github.com/gorilla/mux"
)

var server database.Server

// CreateUser - Register new User(Driver)
func (controller *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, r, http.StatusUnprocessableEntity, err)
	}

	// Create User
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		formattedError := utils.Error(err.Error())
		responses.ERROR(w, r, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, r, http.StatusCreated, "New user successfully created", userCreated)

}

// Get User(Driver) by ID
func (controller *Controller) GetUserByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	userByID, err := user.FindUserByID(server.DB, uint64(userId))
	if err != nil {
		responses.ERROR(w, r, http.StatusInternalServerError, err)
		return
	} else {
		responses.JSON(w, r, http.StatusOK, fmt.Sprintf("Successfully get user with id %d", userId), userByID)
	}

}

// Retrieve all Users(Drivers)
func (controller *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, r, http.StatusInternalServerError, err)
		return
	} else {
		responses.JSON(w, r, http.StatusOK, "Successfully get all users", users)
	}

}
