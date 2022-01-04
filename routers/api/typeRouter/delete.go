package typeRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/classifyType/delete"
	"go-server-template/pkg/apiMap"
	"go-server-template/middleware/global/auth"
)

func DeleteTypeRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_DELETE_TYPE,adminAuth,classifyTypeDelete.DeleteController)
}
