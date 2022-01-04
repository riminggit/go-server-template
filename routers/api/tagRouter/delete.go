package tagRouter

import (
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/tag/delete"
	"github.com/gin-gonic/gin"
)

func DeleteRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_DELETE_TAG,adminAuth,tagDelete.DeleteController)
}
