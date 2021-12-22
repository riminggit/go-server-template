package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/loginAndLayout"
	"go-server-template/pkg/apiMap"
)

func UserLoginRouter(g *gin.RouterGroup) {
	g.POST(apiMap.POST_USER_WX_LOGIN, userLoginAndLayout.UserWXLoginController)
	g.POST(apiMap.POST_USER_LOGIN, userLoginAndLayout.LoginController)
}
