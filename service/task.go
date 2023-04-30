package service

import (
	"context"
	"sync"
	"time"

	"github.com/spf13/cast"

	"to-do-list/pkg/ctl"
	"to-do-list/pkg/util"
	dao2 "to-do-list/repository/db/dao"
	"to-do-list/repository/model"
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
	u, err := dao2.NewUserDao(ctx).FindUserByUserId(userId)
	task := &model.Task{
		User:      u,
		Uid:       u.ID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	err = dao2.NewTaskDao(ctx).CreateTask(task)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (s *TaskSrv) ListTask(ctx context.Context, req *types.ListTasksReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	tasks, total, err := dao2.NewTaskDao(ctx).ListTask(req.Start, req.Limit, u.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	taskRespList := make([]*types.TaskResp, 0)
	for _, v := range tasks {
		taskRespList = append(taskRespList, &types.TaskResp{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(),
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return ctl.RespList(taskRespList, total), nil
}

// ShowTask 展示Task作用
func (s *TaskSrv) ShowTask(ctx context.Context, tId string) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	task, err := dao2.NewTaskDao(ctx).FindTaskByIdAndUserId(u.Id, cast.ToUint(tId))
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	respTask := &types.TaskResp{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		View:      task.View(),
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
	task.AddView() // 增加点击数
	return ctl.RespSuccessWithData(respTask), nil
}

func (s *TaskSrv) DeleteTask(ctx context.Context, tId string) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	err = dao2.NewTaskDao(ctx).DeleteTaskById(u.Id, cast.ToUint(tId))
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	return ctl.RespSuccess(), nil
}

func (s *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskReq, tId string) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	err = dao2.NewTaskDao(ctx).UpdateTask(u.Id, cast.ToUint(tId), req)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (s *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	tasks, err := dao2.NewTaskDao(ctx).SearchTask(u.Id, req.Info)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccessWithData(tasks), nil
}
