package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/todolist-ddd/application/task"
	"github.com/CocaineCong/todolist-ddd/infrastructure/common/log"
	"github.com/CocaineCong/todolist-ddd/infrastructure/interfaces/types"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateTaskReq
		err := ctx.ShouldBind(&req)
		if err == nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		taskEntity, err := types.CreateReqDTO2Entity(ctx, &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "task entity"))
			return
		}
		result, err := task.ServiceImplIns.CreateTask(ctx.Request.Context(), taskEntity)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "failed to create task"))
			return
		}
		resp := types.Entity2TaskResp(result)
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(resp))
	}
}

func ListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListTasksReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		l := task.ServiceImplIns
		resp, err := l.ListTask(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "failed to list task"))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func DetailTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DetailReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		l := task.ServiceImplIns
		result, err := l.DetailTask(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "failed to show task"))
			return
		}
		resp := types.Entity2TaskResp(result)
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(resp))
	}
}

func DeleteTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteTaskReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		l := task.ServiceImplIns
		err = l.DeleteTask(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "failed to delete task"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess())
	}
}

func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(types.UpdateTaskReq)
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		l := task.ServiceImplIns
		t, err := types.UpdateReqDTO2Entity(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "conv failed"))
			return
		}
		err = l.UpdateTask(ctx.Request.Context(), t)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "failed to update task"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess())
	}
}

func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, types.RespError(err, "invalid request"))
			return
		}
		l := task.ServiceImplIns
		resp, err := l.SearchTask(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "failed to search task"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(resp))
	}
}
