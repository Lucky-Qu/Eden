package controller

import (
	"Eden/01_mall_basic/db/model"
	"Eden/01_mall_basic/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateProduct 从表单中获取数据并绑定到dao中定义的模型
func CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "传入数据有误！",
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

// DeleteProduct 从uri中获取删除的id来调用删除方法
func DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
		})
		return
	}
	if err = service.DeleteProduct(idInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err,
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
		data, _ := service.GetProductById(idInt)
		if data == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "没有找到对应资源",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": data,
				"msg":  "更新用户数据成功",
			})
		}
	}
}
func GetProduct(c *gin.Context) {
	queryData := make(map[string]interface{})
	v := c.Request.URL.Query()
	for key, value := range v {
		if len(value) > 0 {
			queryData[key] = value[0]
		}
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
