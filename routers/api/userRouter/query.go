package userRouter

import (
	"go-server-template/pkg/apiMap"
	userQuery "go-server-template/src/user/query"

	"github.com/gin-gonic/gin"
)

func QueryUserRouter(g *gin.RouterGroup) {
	g.POST(apiMap.POST_ADMIN_QUERY_USER_INFO, userQuery.QueryUserDataController)
}
