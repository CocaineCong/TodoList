package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"to-do-list/service"
)

// @Tags USER
// @Summary 用户注册
// @Produce json
// @Accept json
// @Param data body service.UserRegisterService true "用户名, 密码"
// @Success 200 {object} serializer.ResponseUser "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500  {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// @Tags USER
// @Summary 用户登录
// @Produce json
// @Accept json
// @Param     data    body     service.UserLoginService    true      "user_name, password"
// @Success 200 {object} serializer.ResponseUser "{"success":true,"data":{},"msg":"登陆成功"}"
// @Failure 500 {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
