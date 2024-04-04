package services

import (
	"go-ecommerce-api-mssql/config"
	"log"

	"go-ecommerce-api-mssql/models"

	"gorm.io/gorm"
)

var db *gorm.DB = config.DBInstance()

func GetAllProducts() []models.Product {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		log.Fatal(err)
	}
	return products
}

func GetProductById(id string) (*models.Product, *gorm.DB) {
	var product models.Product
	db := db.Where("id = ?", id).Find(&product)
	return &product, db
}
