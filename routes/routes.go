package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"to-do-list/api"
	"to-do-list/middleware"
)

//路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	//middleware.HttpLogToFile(conf.AppMode)
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister) 	//用户注册
		v1.POST("user/login", api.UserLogin)       	//用户登陆
		authed := v1.Group("/")     //需要登陆保护
		authed.Use(middleware.JWT())
		{
			//任务操作
			authed.GET("tasks", api.ListTasks)
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.DELETE("task/:id", api.DeleteTask)
			authed.PUT("task/:id", api.UpdateTask)
			authed.POST("search",api.SearchTasks)
		}
	}
	return r
}
