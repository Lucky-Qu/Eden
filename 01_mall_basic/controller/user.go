package controller

import (
	"Eden/01_mall_basic/db/model"
	"Eden/01_mall_basic/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.UserRegister(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
		"msg":  "注册成功",
	})
}
func UpdateUser(c *gin.Context) {
	var user model.User
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = service.UserChange(idInt, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		userUpdated, _ := service.GetUserById(idInt)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": userUpdated,
			"msg":  "更新成功",
		})
	}
}
func DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = service.UserDelete(idInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "删除成功",
	})
}
func SelectUser(c *gin.Context) {
	queryData := make(map[string]interface{})
	v := c.Request.URL.Query()
	for key, value := range v {
		if len(value) > 0 {
			queryData[key] = value[0]
		}
	}
	if users, err := service.GetUser(&queryData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "出错了",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": users,
			"msg":  "查询成功",
		})
	}
}
