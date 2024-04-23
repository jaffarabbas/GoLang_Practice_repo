package controllers

import (
	"go-ecommerce-api-mssql/services"

	"github.com/gin-gonic/gin"
)

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := services.GetAllProducts() // Assuming 'db' is your GORM database instance
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": products})
	}
}

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("pid")
		product, _ := services.GetProductById(id)
		c.JSON(200, gin.H{"data": product})
	}
}
