package main

import (
	routes "go-jwt-mongo-api/routes"
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

	routes.AuthRouter(router)
	routes.UserRouter(router)

	router.GET("api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "access granted"})
	})

	router.GET("api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "access granted"})
	})

	router.Run(":" + port)
}
