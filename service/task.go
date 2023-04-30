package service

import (
	"context"
	"sync"
	"time"

	"github.com/spf13/cast"

	"to-do-list/pkg/e"
	"to-do-list/pkg/util"
	"to-do-list/repository/dao"
	model2 "to-do-list/repository/model"
	"to-do-list/serializer"
	"to-do-list/types"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (s *TaskSrv) CreateTask(ctx context.Context, req *types.CreateTaskReq, userId uint) (resp interface{}, err error) {
	u, err := dao.NewUserDao(ctx).FindUserByUserId(userId)
	task := &model2.Task{
		User:      u,
		Uid:       u.ID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	code := e.SUCCESS
	err = dao.NewTaskDao(ctx).CreateTask(task)
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}, nil
}

func (s *TaskSrv) ListTask(ctx context.Context, req *types.ListTasksReq, uId uint) (*serializer.Response, error) {
	if req.Limit == 0 {
		req.Limit = 15
	}
	tasks, total, err := dao.NewTaskDao(ctx).ListTask(req.Start, req.Limit, uId)
	if err != nil {
		return nil, err
	}
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total)), nil
}

// ShowTask 展示Task作用
func (s *TaskSrv) ShowTask(ctx context.Context, uId uint, tId string) (resp interface{}, err error) {
	code := e.SUCCESS
	task, err := dao.NewTaskDao(ctx).FindTaskByIdAndUserId(uId, cast.ToUint(tId))
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return nil, err
	}
	task.AddView() // 增加点击数
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}, nil
}

func (s *TaskSrv) DeleteTask(ctx context.Context, uId uint, tId string) (resp interface{}, err error) {
	code := e.SUCCESS
	taskDao := dao.NewTaskDao(ctx)

	err = taskDao.DeleteTaskById(uId, cast.ToUint(tId))
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return nil, err
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}, nil
}

func (s *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskReq, uId uint, tId string) (resp interface{}, err error) {
	err = dao.NewTaskDao(ctx).UpdateTask(uId, cast.ToUint(tId), req)
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}
	code := e.SUCCESS
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "修改成功",
	}, nil
}

func (s *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq, uId uint) (resp interface{}, err error) {
	tasks, err := dao.NewTaskDao(ctx).SearchTask(uId, req.Info)
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}
	code := e.SUCCESS
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildTasks(tasks),
	}, nil
}
