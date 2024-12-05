package dao

import (
	"Eden/01_mall/db/model"
)

func CreateProduct(product *model.Product) error {
	if err := db.Model(&model.Product{}).Create(product); err.Error != nil {
		return db.Error
	}
	return nil
}
func DeleteProduct(num int) error {
	if err := db.Model(&model.Product{}).Where("id = ?", num).Unscoped().Delete(num); err.Error != nil {
		return db.Error
	}
	return nil
}
func UpdateProduct(num int, product *model.Product) error {
	if err := db.Model(&model.Product{}).Where("id = ?", num).Updates(product); err.Error != nil {
		return db.Error
	}
	return nil
}
func GetProduct(queryString string) (*[]model.Product, error) {
	var products []model.Product
	if err := db.Model(&model.Product{}).Where(queryString).Find(&products); err.Error != nil {
		return nil, db.Error
	}
	return &products, nil
}
