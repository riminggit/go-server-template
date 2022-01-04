package classifyRouter

import (
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/classify/update"
	"github.com/gin-gonic/gin"
)

func UpdateClassifyRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_UPDATE_CLASSIFY,adminAuth,classifyUpdate.UpdateClassifyController)
}
