package controller

import (
	"Eden/01_mall/db/model"
	"Eden/01_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定失败",
			"data": nil,
		})
		return
	}
	if err := service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "创建用户失败",
			"data": nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": product,
			"msg":  "创建成功",
		})
	}
}
func DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
			"data": nil,
		})
		return
	}
	if err = service.DeleteProduct(idInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
			"data": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "删除成功",
		})
	}
}
func UpdateProduct(c *gin.Context) {
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
			"data": nil,
		})
		return
	}
	var product model.Product
	if err = c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
			"data": nil,
		})
		return
	}
	if err = service.UpdateProduct(idInt, &product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
			"data": nil,
		})
	} else {
		queryMap := map[string]interface{}{
			"id": idString,
		}
		data, _ := service.GetProduct(&queryMap)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": data,
			"msg":  "更新用户数据成功",
		})
	}
}
func GetProduct(c *gin.Context) {
	queryData := make(map[string]interface{})
	if err := c.ShouldBindJSON(&queryData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
			"data": nil,
		})
		return
	}
	if queryData == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "未输入查询参数",
		})
		return
	} else {
		data, err := service.GetProduct(&queryData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err,
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": data,
			"msg":  "查询成功",
		})
	}
}
