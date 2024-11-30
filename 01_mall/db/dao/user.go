package dao

import (
	"Eden/01_mall/db/model"
	"reflect"
)

// CreateUser 创建新用户
func CreateUser(user *model.User) {
	db.Model(&model.User{}).Create(user)
}

// UpdateUser 更新用户信息
func UpdateUser() {

}

// DeleteUser 删除用户信息
func DeleteUser() {

}

// SelectUser 检索用户信息
func SelectUser(user *model.User) (ResultUser []*model.User) {
	var queryString = ""
	v := reflect.ValueOf(user).Elem()
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsZero() {
			if queryString != "" {
				queryString += "AND"
			}
			queryString = queryString + v.Type().Field(i).Name + " " + "=" + " " + v.Field(i).String()
		}
	}
	db.Where(queryString).Find(&ResultUser)
	return ResultUser
}
