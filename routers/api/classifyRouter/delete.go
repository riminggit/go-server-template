package classifyRouter

import (
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/classify/delete"
	"github.com/gin-gonic/gin"
)

func DeleteClassifyRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_DELETE_CLASSIFY,adminAuth,classifyDelete.DeleteClassifyController)
}
