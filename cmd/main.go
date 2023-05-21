package main

import (
	conf "to-do-list/config"
	"to-do-list/pkg/util"
	"to-do-list/repository/cache"
	"to-do-list/repository/db/dao"
	"to-do-list/routes"
)

// @title ToDoList API
// @version 0.0.1
// @description This is a sample Server pets
// @name FanOne
// @BasePath /api/v1
func main() { // http://localhost:3000/swagger/index.html
	loading()
	// 转载路由 swag init -g common.go
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}

func loading() {
	// 从配置文件读入配置
	conf.Init()
	util.InitLog()
	dao.MySQLInit()
	cache.Redis()
}
