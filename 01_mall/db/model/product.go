package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string  `json:"name"`
	Number int     `json:"number"`
	Price  float64 `json:"price"`
	Info   string  `json:"info"`
}
