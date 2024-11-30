package middleware

import "github.com/gin-gonic/gin"

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
