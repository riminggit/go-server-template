package tagRouter

import (
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/tag/create"
	"github.com/gin-gonic/gin"
)

func CreateRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_CREATE_TAG,adminAuth,tagCreate.CreateController)
}
