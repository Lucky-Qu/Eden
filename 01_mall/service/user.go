package service

import (
	"Eden/01_mall/db/dao"
	"Eden/01_mall/db/model"
	"errors"
	"strconv"
)

func UserRegister(user *model.User) error {
	if err := dao.CreateUser(user); err != nil {
		return err
	}
	return nil
}
func GetUser(queryData *map[string]interface{}) (*[]model.User, error) {
	var queryString string
	for key, value := range *queryData {
		if value != nil {
			if queryString != "" {
				queryString += " && "
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
	user, err := dao.SelectUser(queryString)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func UserChange(id int, userChanged *model.User) error {
	if err := dao.UpdateUser(id, userChanged); err != nil {
		return err
	}
	return nil
}
func UserDelete(id int) error {
	if err := dao.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
func GetUserById(id int) (*model.User, error) {
	var queryString = "id =" + "'" + strconv.Itoa(id) + "'"
	user, err := dao.SelectUser(queryString)
	if err != nil {
		return nil, err
	}
	if len(*user) == 0 {
		return nil, nil
	}
	userUpdated := &(*user)[0]
	return userUpdated, nil
}
