package main

import (
	"go-ecommerce-api-mssql/routes"
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

	routes.ProductRouter(router)

	router.Run(":" + port)
}
