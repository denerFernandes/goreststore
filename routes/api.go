package routes

import (
	"database/sql"

	"github.com/denerFernandes/goreststore/controllers"
	"github.com/denerFernandes/goreststore/database"
	"github.com/denerFernandes/goreststore/middlewares"
	"github.com/gorilla/mux"
)

var (
	db *sql.DB
)

func NewRouter() *mux.Router {

	db = database.Connect()

	r := mux.NewRouter().StrictSlash(true)

	controller := controllers.Controller{}

	r.HandleFunc("/user/signup", controller.Signup(db)).Methods("POST")
	r.HandleFunc("/user/login", controller.Login(db)).Methods("POST")
	r.HandleFunc("/product", middlewares.IsAuth(controller.AddProduct(db))).Methods("POST")

	return r

}
