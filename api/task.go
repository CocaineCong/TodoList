package api

import (
	"github.com/gin-gonic/gin"
	"to-do-list/pkg/util"
	"to-do-list/service"
)

// @Tags TASK
// @Summary 创建任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.CreateTaskService true  "title"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task [post]
func CreateTask(c *gin.Context) {
	createService := service.CreateTaskService{}
	chaim,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// @Tags TASK
// @Summary 获取任务列表
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.ListTasksService true "rush"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /tasks [get]
func ListTasks(c *gin.Context) {
	listService := service.ListTasksService{}
	chaim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// @Tags TASK
// @Summary 展示任务详细信息
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.ShowTaskService true "rush"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task/:id [get]
func ShowTask(c *gin.Context) {
	showTaskService := service.ShowTaskService{}
	res := showTaskService.Show(c.Param("id"))
	c.JSON(200, res)
}

// @Tags TASK
// @Summary 删除任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.DeleteTaskService true "用户信息"
// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task/:id [delete]
func DeleteTask(c *gin.Context) {
	deleteTaskService := service.DeleteTaskService{}
	res := deleteTaskService.Delete(c.Param("id"))
	c.JSON(200, res)
}

// @Tags TASK
// @Summary 修改任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param	data	body	service.DeleteTaskService true "2"
// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task [put]
func UpdateTask(c *gin.Context) {
	updateTaskService := service.UpdateTaskService{}
	if err := c.ShouldBind(&updateTaskService); err == nil {
		res := updateTaskService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// @Tags TASK
// @Summary 查询任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.DeleteTaskService true "2"
// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /search [post]
func SearchTasks(c *gin.Context) {
	searchTaskService := service.SearchTaskService{}
	chaim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.Search(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}