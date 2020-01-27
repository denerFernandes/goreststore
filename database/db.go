package database

import (
	"fmt"
	"log"

	"github.com/denerFernandes/goreststore/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Server - Server struct
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize - Initialize server
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	dbConnection := fmt.Sprintf("host=%s port=%s user=%s sslmode=disable dbname=%s password=%s",
		DbHost, DbPort, DbUser, DbName, DbPassword)

	server.DB, err = gorm.Open(Dbdriver, dbConnection)
	if err != nil {
		fmt.Printf("Cannot connect to %s database ", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.DB.Debug().AutoMigrate(
		&models.User{},
	) //database migration

}
