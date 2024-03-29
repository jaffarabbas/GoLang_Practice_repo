package routes

import (
	controller "go-jwt-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	router.POST("users/register", controller.Register)
	router.POST("user/login", controller.Login)
}
