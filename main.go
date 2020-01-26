package main

import (
	"net/http"
	"os"

	"github.com/denerFernandes/goreststore/routes"
	"github.com/denerFernandes/goreststore/utils"
	"github.com/subosito/gotenv"
)

var (
	_ = gotenv.Load()
)

//Main method
func main() {

	r := routes.NewRouter()

	utils.LogToTerm("Listen on port " + os.Getenv("APP_PORT") + "....")
	utils.LogFatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), routes.LoadCors(r)))
}
