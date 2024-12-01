package controller

import (
	"Eden/01_mall/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	service.UserRegister(c)
}
func UpdateUser(c *gin.Context) {
	service.UserChange(c)
}
func DeleteUser(c *gin.Context) {
	service.UserDelete(c)
}
func SelectUser(c *gin.Context) {
	service.UserGet(c)
}
