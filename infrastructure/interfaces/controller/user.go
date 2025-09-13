package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/todolist-ddd/application/user"
	"github.com/CocaineCong/todolist-ddd/infrastructure/common/ctl"
	"github.com/CocaineCong/todolist-ddd/infrastructure/common/log"
	"github.com/CocaineCong/todolist-ddd/infrastructure/interfaces/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "bind req param failed"))
			return
		}
		userEntity := types.UserReq2Entity(&req)
		resp, err := user.ServiceImplIns.Register(ctx, userEntity)
		if err != nil {
			ctx.JSON(http.StatusOK, ctl.RespError(err, "register failed"))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccessWithData(resp))
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserReq
		err := ctx.ShouldBind(&req)
		if err == nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ctl.RespError(err, "bind req"))
			return
		}
		entity := types.UserReq2Entity(&req)
		resp, err := user.ServiceImplIns.Login(ctx, entity)
		if err != nil {
			ctx.JSON(http.StatusOK, ctl.RespError(err, "login failed"))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccessWithData(resp))
	}
}
