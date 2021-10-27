package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/login"
)

func UserLoginRouter(g *gin.RouterGroup) {
	g.POST("/wxLogin", userLogin.UserWXLoginController)
}
