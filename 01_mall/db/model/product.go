package model

import "gorm.io/gorm"

// Product 商品的属性
type Product struct {
	gorm.Model
	Name  string
	Price int
	Stock int
	Info  string
}
