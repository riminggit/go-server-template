package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/loginAndExit"
	"go-server-template/pkg/apiMap"
)

func UserLoginRouter(g *gin.RouterGroup) {
	g.POST(apiMap.POST_USER_WX_LOGIN, userloginAndExit.UserWXLoginController)
	g.POST(apiMap.POST_USER_LOGIN, userloginAndExit.LoginController)
}
