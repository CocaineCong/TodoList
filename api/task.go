package api

import (
	"github.com/gin-gonic/gin"
	"to-do-list/pkg/logging"
	"to-do-list/pkg/util"
	"to-do-list/service"
)

func CreateTask(c *gin.Context) {
	createService := service.CreateTaskService{}
	chaim,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func ListTasks(c *gin.Context) {
	listService := service.ListTasksService{}
	chaim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func ShowTask(c *gin.Context) {
	showTaskService := service.ShowTaskService{}
	res := showTaskService.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteTask(c *gin.Context) {
	deleteTaskService := service.DeleteTaskService{}
	res := deleteTaskService.Delete(c.Param("id"))
	c.JSON(200, res)
}

func UpdateTask(c *gin.Context) {
	updateTaskService := service.UpdateTaskService{}
	if err := c.ShouldBind(&updateTaskService); err == nil {
		res := updateTaskService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func SearchTasks(c *gin.Context) {
	searchTaskService := service.SearchTaskService{}
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.Search()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
