package main

import "Eden/01_mall_basic/routes"

func main() {
	//运行服务器
	routes.Run(routes.NewRouter())
}
