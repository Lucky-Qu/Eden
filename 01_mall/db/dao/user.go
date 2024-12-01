package dao

import (
	"Eden/01_mall/db/model"
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
func SelectUser(username string, sex string) (ResultUser []*model.User) {
	queryString := ""
	if len(username) > 0 {
		queryString = queryString + "username = " + "\"" + username + "\""
		if len(sex) > 0 {
			queryString = queryString + " && " + "sex = " + "\"" + sex + "\""
		}
	} else if len(sex) > 0 {
		queryString = queryString + "sex = " + "\"" + sex + "\""
	}
	db.Model(&model.User{}).Where(queryString).Find(&ResultUser)
	return ResultUser
}
