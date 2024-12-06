package dao

import (
	"Eden/01_mall_basic/db/model"
)

func CreateProduct(product *model.Product) error {
	if err := db.Model(&model.Product{}).Create(product).Error; err != nil {
		return err
	}
	return nil
}
func DeleteProduct(num int) error {
	if err := db.Model(&model.Product{}).Where("id = ?", num).Unscoped().Delete(num).Error; err != nil {
		return err
	}
	return nil
}
func UpdateProduct(id int, product *model.Product) error {
	if err := db.Model(&model.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}
	return nil
}
func GetProduct(queryString string) (*[]model.Product, error) {
	var products []model.Product
	if err := db.Model(&model.Product{}).Where(queryString).Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}
