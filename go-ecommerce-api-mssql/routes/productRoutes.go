package routes

import (
	"go-ecommerce-api-mssql/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.Engine) {
	// router.Use(middleware.AuthMiddleware())
	router.GET("products", controllers.GetProducts())
	router.GET("products/:pid", controllers.GetProduct())
}
