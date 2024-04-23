package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func DBInstance() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DbUrl := os.Getenv("DATA_BASE_URL")
	db, err := gorm.Open(sqlserver.Open(DbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
