package mysql

import (
	"fmt"
	"os"
	"database/sql"
	
	_ "github.com/lib/pq"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

var DB *sql.DB

// Connection Database
func DatabaseInit() {
	var err error
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_PORT = os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}