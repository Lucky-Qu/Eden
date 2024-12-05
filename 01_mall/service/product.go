package service

import (
	"Eden/01_mall/db/dao"
	"Eden/01_mall/db/model"
)

func CreateProduct(product *model.Product) error {
	if err := dao.CreateProduct(product); err != nil {
		return err
	}
	return nil
}
func DeleteProduct(id int) error {
	if err := dao.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}
func UpdateProduct(id int, product *model.Product) error {
	if err := dao.UpdateProduct(id, product); err != nil {
		return err
	}
	return nil
}
func GetProduct(queryData *map[string]interface{}) (*[]model.Product, error) {
	var queryString = ""
	//查询字符串的拼接

	products, err := dao.GetProduct(queryString)
	if err != nil {
		return nil, err
	}
	return products, nil
}
