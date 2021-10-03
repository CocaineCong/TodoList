package main

import (
	"to-do-list/conf"
	"to-do-list/routes"
)

func main() {
	//从配置文件读入配置
	conf.Init()
	//转载路由
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}