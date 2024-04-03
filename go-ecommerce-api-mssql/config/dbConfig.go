package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DBInstance() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DbUrl := os.Getenv("DATA_BASE_URL")
	db, err := sql.Open("sqlserver", DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var db *sql.DB = DBInstance()

func GetDB() *sql.DB {
	return db
}
