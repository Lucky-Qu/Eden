package model

import "gorm.io/gorm"

// Product 商品的属性
type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
	Info  string  `json:"info"`
}
