package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/user/query"
)

func QueryUserRouter(g *gin.RouterGroup) {
	g.POST("/query/user-info", userQuery.QueryUserDataController)
}
