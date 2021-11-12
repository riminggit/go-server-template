package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/src/loginAndLayout"
)

func UserLoginRouter(g *gin.RouterGroup) {
	g.GET("/layout", JWTMiddleware.JWT(), userLoginAndLayout.LayoutController)
	g.POST("/wxLogin", userLoginAndLayout.UserWXLoginController)
	g.POST("/login", userLoginAndLayout.LoginController)
}
