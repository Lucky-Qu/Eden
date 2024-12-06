package middleware

import "github.com/gin-gonic/gin"

// Middleware 中间件
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
