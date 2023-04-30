package service

import (
	"context"
	"sync"
	"time"

	"github.com/spf13/cast"

	"to-do-list/pkg/ctl"
	"to-do-list/pkg/util"
	"to-do-list/repository/dao"
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
	u, err := dao.NewUserDao(ctx).FindUserByUserId(userId)
	task := &model.Task{
		User:      u,
		Uid:       u.ID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	err = dao.NewTaskDao(ctx).CreateTask(task)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}

func (s *TaskSrv) ListTask(ctx context.Context, req *types.ListTasksReq, uId uint) (resp interface{}, err error) {
	if req.Limit == 0 {
		req.Limit = 15
	}
	tasks, total, err := dao.NewTaskDao(ctx).ListTask(req.Start, req.Limit, uId)
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
func (s *TaskSrv) ShowTask(ctx context.Context, uId uint, tId string) (resp interface{}, err error) {
	task, err := dao.NewTaskDao(ctx).FindTaskByIdAndUserId(uId, cast.ToUint(tId))
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
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

func (s *TaskSrv) DeleteTask(ctx context.Context, uId uint, tId string) (resp interface{}, err error) {
	err = dao.NewTaskDao(ctx).DeleteTaskById(uId, cast.ToUint(tId))
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}

	return ctl.RespSuccess(), nil
}

func (s *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskReq, uId uint, tId string) (resp interface{}, err error) {
	err = dao.NewTaskDao(ctx).UpdateTask(uId, cast.ToUint(tId), req)
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}
	return ctl.RespSuccess(), nil
}

func (s *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq, uId uint) (resp interface{}, err error) {
	tasks, err := dao.NewTaskDao(ctx).SearchTask(uId, req.Info)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccessWithData(tasks), nil
}
