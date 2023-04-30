package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"to-do-list/api"
	_ "to-do-list/docs" // 这里需要引入本地已生成文档
	"to-do-list/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default() // 生成了一个WSGI应用程序实例
	store := cookie.NewStore([]byte("something-very-secret"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 开启swag
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())
		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			// 任务操作
			authed.GET("tasks", api.ListTasksHandler())
			authed.POST("task", api.CreateTaskHandler())
			authed.GET("task/:id", api.ShowTaskHandler())
			authed.DELETE("task/:id", api.DeleteTaskHandler())
			authed.PUT("task/:id", api.UpdateTaskHandler())
			authed.POST("search", api.SearchTasksHandler())
			// Tip:其实
		}
	}
	return r
}
