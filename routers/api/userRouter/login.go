package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/loginAndLayout"
)

func UserLoginRouter(g *gin.RouterGroup) {
	g.POST("/wxLogin", userLoginAndLayout.UserWXLoginController)
	g.POST("/login", userLoginAndLayout.LoginController)
}
