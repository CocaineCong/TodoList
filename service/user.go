package service

import (
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"to-do-list/model"
	"to-do-list/pkg/e"
	"to-do-list/pkg/util"
	"to-do-list/serializer"
)

//UserRegisterService 用户注册服务
type UserService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password  string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}


func (service *UserService) Valid(userId, status interface{}) *serializer.Response {
	var code int
	count := 0
	err := model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).Count(&count).Error
	if err != nil {
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if count > 0 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return nil
}


func (service *UserService) Register(userID, status interface{}) *serializer.Response {
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
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
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
func (service *UserService)Login() serializer.Response {
	var user model.User
	code := e.SUCCESS
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID,service.UserName, service.Password, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
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