package service

import (
	"context"
	"errors"
	"sync"

	"gorm.io/gorm"

	"to-do-list/pkg/ctl"
	"to-do-list/pkg/util"
	"to-do-list/repository/db/dao"
	"to-do-list/repository/db/model"
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
	u, err := userDao.FindUserByUserName(req.UserName)
	switch err { // TODO 优化一下
	case gorm.ErrRecordNotFound:
		u = &model.User{
			UserName: req.UserName,
		}
		// 密码加密存储
		if err = u.SetPassword(req.Password); err != nil {
			util.LogrusObj.Info(err)
			return
		}

		if err = userDao.CreateUser(u); err != nil {
			util.LogrusObj.Info(err)
			return
		}

		return ctl.RespSuccess(), nil
	case nil:
		err = errors.New("用户已存在")
		return
	default:
		return
	}
}

// Login 用户登陆函数
func (s *UserSrv) Login(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	if err == gorm.ErrRecordNotFound {
		err = errors.New("用户不存在")
		return
	}

	if !user.CheckPassword(req.Password) {
		err = errors.New("账号/密码错误")
		util.LogrusObj.Info(err)
		return
	}

	token, err := util.GenerateToken(user.ID, req.UserName, 0)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	u := &types.UserResp{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
	uResp := &types.TokenData{
		User:  u,
		Token: token,
	}

	return ctl.RespSuccessWithData(uResp), nil
}
