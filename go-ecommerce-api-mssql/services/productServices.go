package services

import (
	"go-ecommerce-api-mssql/config"

	"go-ecommerce-api-mssql/models"

	"gorm.io/gorm"
)

var db *gorm.DB = config.DBInstance()

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(id string) (*models.Product, *gorm.DB) {
	var product models.Product
	db := db.Where("id = ?", id).Find(&product)
	return &product, db
}
