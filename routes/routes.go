package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"to-do-list/api"
	"to-do-list/conf"
	_ "to-do-list/docs" // 这里需要引入本地已生成文档
	"to-do-list/middleware"
	"to-do-list/pkg/log"
)


//路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()   //生成了一个WSGI应用程序实例
	store := cookie.NewStore([]byte("something-very-secret"))
	log.HttpLogToFile(conf.AppMode) 	// 日志输出
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))	// 开启swag
	r.Use(sessions.Sessions("mysession", store))
	r.Use(Recovery)
	v1 := r.Group("api/v1")
	{
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
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

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logging.Error("gin catch error: ", r)
			c.JSON(http.StatusInternalServerError,"系统内部错误")
		}
	}()
	c.Next()
}
