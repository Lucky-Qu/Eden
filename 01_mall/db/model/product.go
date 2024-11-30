package model

import "gorm.io/gorm"

// Product 商品的属性
type Product struct {
	gorm.Model
	Name   string  `json:"name"`
	Number int     `json:"number"`
	Price  float64 `json:"price"`
	Info   string  `json:"info"`
}
