package main

import "Eden/01_mall/routes"

func main() {
	//运行服务器
	routes.Run(routes.NewRouter())
}
