package service

import (
	"Eden/01_mall/db/dao"
	"Eden/01_mall/db/model"
	"github.com/gin-gonic/gin"
	"strconv"
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
func UserGet(c *gin.Context) {
	var username string
	var sex string
	username, _ = c.GetQuery("username")
	sex, _ = c.GetQuery("sex")
	result := dao.SelectUser(username, sex)
	if result != nil {
		c.JSON(200, gin.H{
			"msg":  "查询成功",
			"user": result,
		})
	} else {
		c.JSON(400, gin.H{
			"msg": "未查询到用户",
		})
	}
}
func UserChange(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	var user model.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	userBefore := user
	dao.UpdateUser(id, &user)
	c.JSON(200, gin.H{
		"msg":        "更改成功",
		"userBefore": userBefore,
		"userAfter":  user,
	})
}
func UserDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	dao.DeleteUser(id)
	c.JSON(200, gin.H{
		"msg": "删除成功",
	})
}
