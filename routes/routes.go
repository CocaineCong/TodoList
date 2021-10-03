package routes

import (
	"to-do-list/api"
	"to-do-list/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister) //用户注册
		v1.POST("user/login", api.UserLogin)       //用户登陆
		//商品操作
		v1.GET("products", api.ListProducts)
		v1.POST("products_same", api.ListSameProducts)
		v1.GET("products/:id", api.ShowProduct)
		v1.GET("categories", api.ListCategories)    //商品分类操作

		//v1.POST("payments",api.InitPay)

		authed := v1.Group("/")            //需要登陆保护
		authed.Use(middleware.JWT())
		{
			authed.PUT("products", api.UpdateProduct)
			authed.GET("SearchProducts",api.SearchProducts)
			authed.POST("products",api.CreateProduct)   //创建商品
			authed.GET("ping", api.CheckToken)                //验证token
		}
	}
	return r
}
