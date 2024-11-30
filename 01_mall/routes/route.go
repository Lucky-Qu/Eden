package routes

import (
	"Eden/01_mall/config"
	"Eden/01_mall/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 新建一个路由
func NewRouter() *gin.Engine {
	g := gin.Default()
	g.Use(middleware.Middleware())
	return g
}

// Run 运行并监听端口
func Run(g *gin.Engine) {
	err := g.Run(config.ListenLocation)
	if err != nil {
		panic(err)
	}
}
