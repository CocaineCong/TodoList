package main

import (
	"to-do-list/conf"
	"to-do-list/routes"
)

// @title ToDoList API
// @version 0.0.1
// @description This is a sample Server pets
// @name FanOne
// @BasePath /api/v1
func main() { // http://localhost:3000/swagger/index.html
	//从配置文件读入配置
	conf.Init()
	//转载路由 swag init -g common.go
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}