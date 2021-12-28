package JWTMiddleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-server-template/config"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/redis"
	"go-server-template/pkg/utils"
	"net/http"
)

// 需要校验的路由用
// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := app.GetHeaderToken(c)
		// token := c.Query("token") // 获取链接上的参数
		if token == "" || token == "undefined" {
			token = c.Query("token")
			if token == "" || token == "undefined" {
				code = e.USER_NOT_LOGIN
			} else {
				code = verifyType(token)
			}
		} else {
			code = verifyType(token)
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

func verifyType(token string) int {

	var code int
	code = e.SUCCESS
	if projectConfig.AppConfig.BaseConfig.JWT_VERIFY_TYPE == "token" {
		// 解析token校验
		_, err := util.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				code = e.USER_NOT_LOGIN
			}
		}
	} else if projectConfig.AppConfig.BaseConfig.JWT_VERIFY_TYPE == "redis" {
		// 使用redis校验token
		userInfo := Redis.GetValue(token)
		if userInfo == "" {
			code = e.USER_NOT_LOGIN
		}
	}
	return code
}
