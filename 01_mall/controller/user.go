package controller

import (
	"Eden/01_mall/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	service.UserRegister(c)
}
func UpdateUser(c *gin.Context) {

}
func DeleteUser(c *gin.Context) {

}
func SelectUser(c *gin.Context) {

}
