package model

import "gorm.io/gorm"

// User 用户的属性
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
