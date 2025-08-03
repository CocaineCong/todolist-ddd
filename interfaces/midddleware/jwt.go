package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	lctx "github.com/CocaineCong/todolist-ddd/infra/common/context"
	"github.com/CocaineCong/todolist-ddd/infra/common/e"
	"github.com/CocaineCong/todolist-ddd/infra/common/jwt"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "缺少Token",
			})
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}

		if code != e.SUCCESS {
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
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
