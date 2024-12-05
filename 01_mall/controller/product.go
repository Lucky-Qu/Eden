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
			"msg":  err,
			"data": nil,
		})
		return
	}
	if err := service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
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

}
func GetProduct(c *gin.Context) {

}
