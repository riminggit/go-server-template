package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/user/query"
	"go-server-template/pkg/apiMap"
)

func QueryUserRouter(g *gin.RouterGroup) {
	g.POST(apiMap.POST_USER_QUERY_USER_INFO, userQuery.QueryUserDataController)
}
