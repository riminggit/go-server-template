package tagRouter

import (
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/tag/update"
	"github.com/gin-gonic/gin"
)

func UpdateRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_UPDATE_TAG,adminAuth,tagUpdate.UpdateController)
}
