package service

import (
	"Eden/01_mall/db/dao"
	"Eden/01_mall/db/model"
	"errors"
	"strconv"
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
	for key, value := range *queryData {
		if value != nil {
			if queryString != "" {
				queryString += "&&"
			}
			switch value.(type) {
			case string:
				queryString += key + " = " + "'" + value.(string) + "'"
			case int:
				queryString += key + " = " + "'" + strconv.Itoa(value.(int)) + "'"
			default:
				return nil, errors.New("传入了错误类型的参数")
			}
		}
	}
	products, err := dao.GetProduct(queryString)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(id int) (*model.Product, error) {
	var queryString = "id = " + "'" + strconv.Itoa(id) + "'"
	products, err := dao.GetProduct(queryString)
	if err != nil {
		return nil, err
	}
	if len(*products) == 0 {
		return nil, nil
	}
	product := &(*products)[0]
	return product, nil
}
