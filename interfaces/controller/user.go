package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/todolist-ddd/application/user"
	"github.com/CocaineCong/todolist-ddd/infra/common/ctl"
	"github.com/CocaineCong/todolist-ddd/infra/common/util"
	"github.com/CocaineCong/todolist-ddd/interfaces/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserReq
		if err := ctx.ShouldBind(&req); err == nil {
			resp, err := user.ServiceImplIns.Register(ctx, &req)
			if err != nil {
				ctx.JSON(http.StatusOK, ctl.RespError(err, "register failed"))
				return
			}
			ctx.JSON(http.StatusOK, ctl.RespSuccessWithData(resp))
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "bind req param failed"))
		}
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserReq
		if err := ctx.ShouldBind(&req); err == nil {
			resp, err := user.ServiceImplIns.Login(ctx, &req)
			if err != nil {
				ctx.JSON(http.StatusOK, ctl.RespError(err, "login failed"))
				return
			}
			ctx.JSON(http.StatusOK, ctl.RespSuccessWithData(resp))
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "bind req"))
		}
	}
}
