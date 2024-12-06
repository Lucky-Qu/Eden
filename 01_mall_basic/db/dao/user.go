package dao

import (
	"Eden/01_mall_basic/db/model"
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
	if err := db.Unscoped().Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// SelectUser 检索用户信息
func SelectUser(queryString string) (*[]model.User, error) {
	var users []model.User
	if err := db.Where(queryString).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
