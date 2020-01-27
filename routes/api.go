package routes

import (
	"github.com/denerFernandes/goreststore/controllers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	controller := controllers.Controller{}

	r.HandleFunc("/user/register", controller.CreateUser).Methods("POST")
	r.HandleFunc("/user/login", controller.Login).Methods("POST")
	//r.HandleFunc("/product", middlewares.IsAuth(controller.AddProduct(db))).Methods("POST")

	return r

}
