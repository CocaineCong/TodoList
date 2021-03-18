package main

import (
	"cmall/conf"
	"cmall/routes"
)

func main() {
	//从配置文件读入配置
	conf.Init()
	//转载路由
	r := routes.NewRouter()
	_ = r.Run(":3000")
}