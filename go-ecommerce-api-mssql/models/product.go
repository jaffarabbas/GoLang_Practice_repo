package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Pid           int     `json:"pid"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	Image         string  `json:"image"`
	Created_On    string  `json:"created_on"`
	Cid           int     `json:"cid"`
	ProductStatus string  `json:"product_status"`
}
