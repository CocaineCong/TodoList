package service

import (
	"context"
	"errors"
	"sync"

	"gorm.io/gorm"

	"to-do-list/pkg/e"
	"to-do-list/pkg/util"
	"to-do-list/repository/dao"
	"to-do-list/repository/model"
	"to-do-list/serializer"
	"to-do-list/types"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) Register(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound:
		user = &model.User{
			UserName: req.UserName,
		}
		// 密码加密存储
		if err = user.SetPassword(req.Password); err != nil {
			util.LogrusObj.Info(err)
			return nil, err
		}

		if err = userDao.CreateUser(user); err != nil {
			util.LogrusObj.Info(err)
			return nil, err
		}

		return serializer.Response{
			Status: e.SUCCESS,
			Msg:    e.GetMsg(e.SUCCESS),
		}, nil
	case nil:
		return nil, errors.New("用户已存在")
	default:
		return nil, err
	}
}

// Login 用户登陆函数
func (s *UserSrv) Login(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("用户不存在，请注册！")
	}
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("账号/密码错误")
	}
	token, err := util.GenerateToken(user.ID, req.UserName, 0)
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}

	return &serializer.Response{
		Data: types.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}, nil
}
