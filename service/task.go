package service

import (
	"time"
	"to-do-list/model"
	"to-do-list/pkg/e"
	"to-do-list/pkg/util"
	"to-do-list/serializer"
)

//展示任务详情的服务
type ShowTaskService struct {
}

//删除任务的服务
type DeleteTaskService struct {
}

//更新任务的服务
type UpdateTaskService struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

//创建任务的服务
type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

//搜索任务的服务
type SearchTaskService struct {
	Info string `form:"info" json:"info"`
}

type ListTasksService struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	code := e.SUCCESS
	err := model.DB.Create(&task).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListTasksService) List(id uint) serializer.Response {
	var tasks []model.Task
	var total int64
	if service.Limit == 0 {
		service.Limit = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid = ?", id).Count(&total).
		Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))
}

func (service *ShowTaskService) Show(id string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	err := model.DB.First(&task, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	task.AddView() //增加点击数
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	err := model.DB.First(&task, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&task).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *UpdateTaskService) Update(id string) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id = ?", id).First(&task)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	code := e.SUCCESS
	err := model.DB.Save(&task).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "修改成功",
	}
}

func (service *SearchTaskService) Search(uId uint) serializer.Response {
	var tasks []model.Task
	code := e.SUCCESS
	model.DB.Where("uid=?", uId).Preload("User").First(&tasks)
	err := model.DB.Model(&model.Task{}).Where("title LIKE ? OR content LIKE ?",
		"%"+service.Info+"%", "%"+service.Info+"%").Find(&tasks).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildTasks(tasks),
	}
}
