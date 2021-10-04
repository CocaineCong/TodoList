package service

import (
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"to-do-list/model"
	"to-do-list/pkg/e"
	"to-do-list/pkg/util"
	"to-do-list/serializer"
)

//UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password  string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

//UserLoginService 管理用户登陆的服务
type UserLoginService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password  string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

//valid 验证表单 验证用户是否存在
func (service *UserRegisterService) Valid(userId, status interface{}) *serializer.Response {
	var code int
	count := 0
	err := model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).Count(&count).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if count > 0 {
		code = e.ERROR_EXIST_USER
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return nil
}

//Register 用户注册
func (service *UserRegisterService) Register(userID, status interface{}) *serializer.Response {
	user := model.User{
		UserName: service.UserName,
	}
	code := e.SUCCESS
	//表单验证
	if res := service.Valid(userID, status); res != nil {
		return res
	}
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ERROR_FAIL_ENCRYPTION
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//Login 用户登陆函数
func (service *UserLoginService) Login(userID, status interface{}) serializer.Response {
	var user model.User
	code := e.SUCCESS
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ERROR_NOT_EXIST_USER
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = e.ERROR_NOT_COMPARE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID,service.UserName, service.Password, 0)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}