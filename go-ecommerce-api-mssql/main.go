package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())

	dbConnection, err := config.DBInstance()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()
}
