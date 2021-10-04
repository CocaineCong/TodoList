package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"to-do-list/service"
)

//UserRegister 用户注册
func UserRegister(c *gin.Context) {
	session := sessions.Default(c)
	status := 200
	userID := session.Get("userId")
	var userRegisterService service.UserRegisterService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(userID, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//UserLogin 用户登陆接口
func UserLogin(c *gin.Context) {
	session := sessions.Default(c)
	status := 200
	userID := session.Get("userId")
	var userLoginService service.UserLoginService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(userID, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
