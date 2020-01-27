package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/denerFernandes/goreststore/database"
	"github.com/denerFernandes/goreststore/routes"
	"github.com/denerFernandes/goreststore/utils"
	"github.com/joho/godotenv"
)

var server = database.Server{}

//Main method
func main() {

	// Loading .env file
	var err error
	err = godotenv.Load()

	// Checking error for loading .env file
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// Initialize Database and Server
	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	// Run server
	utils.LogToTerm("Listen on port " + os.Getenv("APP_PORT") + "....")

	r := routes.NewRouter()
	fmt.Println("Listening to port " + os.Getenv("APP_PORT"))
	utils.LogFatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), routes.LoadCors(r)))

}
