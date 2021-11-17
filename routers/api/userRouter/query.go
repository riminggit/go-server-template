package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/user"
)

func QueryUserRouter(g *gin.RouterGroup) {
	g.POST("/query/user-info", user.QueryUserDataController)
}
