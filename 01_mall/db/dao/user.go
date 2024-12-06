package dao

import (
	"Eden/01_mall/db/model"
)

// CreateUser 创建新用户
func CreateUser(user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser 更新用户信息
func UpdateUser(id int, user *model.User) error {
	if err := db.Model(&model.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser 删除用户信息
func DeleteUser(id int) error {
	if err := db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// SelectUser 检索用户信息
func SelectUser(username string, sex string) ([]*model.User, error) {
	queryString := ""
	if len(username) > 0 {
		queryString = queryString + "username = " + "\"" + username + "\""
		if len(sex) > 0 {
			queryString = queryString + " && " + "sex = " + "\"" + sex + "\""
		}
	} else if len(sex) > 0 {
		queryString = queryString + "sex = " + "\"" + sex + "\""
	}
	var ResultUser []*model.User
	if err := db.Model(&model.User{}).Where(queryString).Find(&ResultUser).Error; err != nil {
		return nil, err
	}
	return ResultUser, nil
}
