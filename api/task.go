package api

import (
	"github.com/gin-gonic/gin"
	"to-do-list/pkg/logging"
	"to-do-list/pkg/util"
	"to-do-list/service"
)

func CreateTask(c *gin.Context) {
	service := service.CreateTaskService{}
	chaim,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func ListTasks(c *gin.Context) {
	service := service.ListTasksService{}
	chaim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func ShowTask(c *gin.Context) {
	service := service.ShowTaskService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteTask(c *gin.Context) {
	service := service.DeleteTaskService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateTask(c *gin.Context) {
	service := service.UpdateTaskService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func SearchTasks(c *gin.Context) {
	service := service.SearchTaskService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Search()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
