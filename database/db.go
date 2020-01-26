package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/denerFernandes/goreststore/utils"
	_ "github.com/lib/pq"
)

var db *sql.DB

// ConnectDB - Connect in database and return the connection
func Connect() *sql.DB {

	pgURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))

	db, err := sql.Open("postgres", pgURL)
	utils.LogFatal(err)

	err = db.Ping()
	utils.LogFatal(err)

	return db
}
