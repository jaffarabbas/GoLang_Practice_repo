package controllers

import (
	"go-ecommerce-api-mssql/services"

	"github.com/gin-gonic/gin"
)

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		product := services.GetAllProducts()
		c.JSON(200, gin.H{"data": product})
	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("pid")
		product, _ := services.GetProductById(id)
		c.JSON(200, gin.H{"data": product})
	}
}
