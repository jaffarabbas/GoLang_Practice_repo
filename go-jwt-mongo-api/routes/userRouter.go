package routes

import (
	controller "go-jwt-mongo-api/controllers"
	"go-jwt-mongo-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.GET("users", controller.GetUsers())
	router.GET("users/:user_id", controller.GetUser())
}
