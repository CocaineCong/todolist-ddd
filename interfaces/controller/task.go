package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/todolist-ddd/application/task"
	"github.com/CocaineCong/todolist-ddd/infra/common/ctl"
	"github.com/CocaineCong/todolist-ddd/infra/common/util"
	"github.com/CocaineCong/todolist-ddd/interfaces/types"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			resp, err := task.ServiceImplIns.CreateTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusOK, ctl.RespError(err, "failed to create task"))
				return
			}
			ctx.JSON(http.StatusOK, ctl.RespSuccessWithData(resp))
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "invalid request"))
		}

	}
}

func ListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListTasksReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := task.ServiceImplIns
			resp, err := l.ListTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusOK, ctl.RespError(err, "failed to list task"))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "invalid request"))
		}

	}
}

func DetailTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DetailReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := task.ServiceImplIns
			resp, err := l.DetailTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusOK, ctl.RespError(err, "failed to show task"))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "invalid request"))
		}

	}
}

func DeleteTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := task.ServiceImplIns
			resp, err := l.DeleteTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusOK, ctl.RespError(err, "failed to delete task"))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "invalid request"))
		}

	}
}

func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(types.UpdateTaskReq)
		if err := ctx.ShouldBind(&req); err == nil {
			l := task.ServiceImplIns
			err = l.UpdateTask(ctx.Request.Context(), req)
			if err != nil {
				ctx.JSON(http.StatusOK, ctl.RespError(err, "failed to update task"))
				return
			}
			ctx.JSON(http.StatusOK, ctl.RespSuccess())
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "invalid request"))
		}

	}
}

func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := task.ServiceImplIns
			resp, err := l.SearchTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusOK, ctl.RespError(err, "failed to search task"))
				return
			}
			ctx.JSON(http.StatusOK, ctl.RespSuccessWithData(resp))
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "invalid request"))
		}

	}
}
