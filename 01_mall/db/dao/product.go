package dao

import "Eden/01_mall/db/model"

func CreateProduct(product *model.Product) {
	db.Model(&model.Product{}).Create(product)
}
func DeleteProduct(num int) {
	db.Model(&model.Product{}).Where("id = ?", num).Unscoped().Delete(num)
}
func UpdateProduct(num int, product *model.Product) {
	db.Model(&model.Product{}).Where("id = ?", num).Updates(product)
}
func GetProduct(queryString string) []model.Product {
	var products []model.Product
	db.Model(&model.Product{}).Where(queryString).Find(&products)
	return products
}
