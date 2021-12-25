package authMiddleware

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/utils"
	"go-server-template/pkg/e"
	"net/http"
)

// 校验是否为管理员
func UserAuth() gin.HandlerFunc {

	//返回中间件
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS

		//中间件逻辑
		userInfoRes := util.GetUserInfo(c)
		if userInfoRes.Code != e.SUCCESS {
			code = userInfoRes.Code
		}

		if userInfoRes.Data.IsAdmin != 1 {
			code = e.UN_AUTHORIZED_NOT_ADMIN
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
