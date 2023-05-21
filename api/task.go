package api

import (
	"net/http"

	"to-do-list/consts"
	"to-do-list/pkg/util"
	"to-do-list/service"
	"to-do-list/types"

	"github.com/gin-gonic/gin"
)

// CreateTaskHandler @Tags TASK
// @Summary 创建任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.CreateTaskService true  "title"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task [post]
func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.CreateTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

// ListTaskHandler @Tags TASK
// @Summary 获取任务列表
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.ListTasksService true "rush"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /tasks [get]
func ListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListTasksReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.Limit == 0 {
				req.Limit = consts.BasePageLimit
			}
			l := service.GetTaskSrv()
			resp, err := l.ListTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

// ShowTaskHandler @Tags TASK
// @Summary 展示任务详细信息
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.ShowTaskService true "rush"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task/:id [get]
func ShowTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ShowTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.ShowTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

// DeleteTaskHandler @Tags TASK
// @Summary 删除任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.DeleteTaskService true "用户信息"
// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task/:id [delete]
func DeleteTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.DeleteTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

// UpdateTaskHandler @Tags TASK
// @Summary 修改任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param	data	body	service.DeleteTaskService true "2"
// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task [put]
func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(types.UpdateTaskReq)
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.UpdateTask(ctx.Request.Context(), req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}

// SearchTaskHandler @Tags TASK
// @Summary 查询任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.DeleteTaskService true "2"
// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /search [post]
func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetTaskSrv()
			resp, err := l.SearchTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}
}
