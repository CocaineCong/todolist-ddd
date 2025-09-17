package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/todolist-ddd/infrastructure/auth"
	lctx "github.com/CocaineCong/todolist-ddd/infrastructure/common/context"
	"github.com/CocaineCong/todolist-ddd/infrastructure/consts"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = consts.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(consts.InvalidParams, gin.H{
				"status": code,
				"msg":    consts.GetMsg(code),
				"data":   "缺少Token",
			})
			c.Abort()
			return
		}
		jwtService := auth.NewJWTTokenService()
		claims, err := jwtService.ParseToken(c.Request.Context(), token)
		if err != nil {
			code = consts.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = consts.ErrorAuthCheckTokenTimeout
		}

		if code != consts.SUCCESS {
			c.JSON(consts.InvalidParams, gin.H{
				"status": code,
				"msg":    consts.GetMsg(code),
				"data":   "可能是身份过期了，请重新登录",
			})
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(
			lctx.NewContext(
				c.Request.Context(),
				&lctx.UserInfo{Id: claims.Id, Name: claims.Username},
			),
		)
		c.Next()
	}
}
