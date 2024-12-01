package service

import (
	"Eden/01_mall/db/dao"
	"Eden/01_mall/db/model"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	dao.CreateUser(&user)
	c.JSON(200, gin.H{
		"msg":  "创建成功",
		"user": user,
	})
}
