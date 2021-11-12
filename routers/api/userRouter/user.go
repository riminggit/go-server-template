package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/user"
)

func UserRelevantRouter(g *gin.RouterGroup) {
	g.POST("/query/userInfo", user.QueryUserDataController)
}
