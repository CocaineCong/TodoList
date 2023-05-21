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
			authed.POST("task_create", api.CreateTaskHandler())
			authed.GET("task_list", api.ListTaskHandler())
			authed.GET("task_show", api.ShowTaskHandler())
			authed.POST("task_update", api.UpdateTaskHandler())
			authed.POST("task_search", api.SearchTaskHandler())
			authed.POST("task_delete", api.DeleteTaskHandler())
			// Tip: 这个RESTful api的路由其实一般不怎么用，RESTful风格的路由算 模糊路由
			// 但是在企业实际生产中，我们一般需要 精确路由 给到网关进行对应的配置，所以我们的请求一般都是清一色的GET、POST
			// 虽然RESTful不怎么用，但是也是要知道，也要会使用这个风格。
		}
	}
	return r
}
