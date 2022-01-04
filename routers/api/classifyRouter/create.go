package classifyRouter

import (
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/classify/create"
	"github.com/gin-gonic/gin"
)

func CreateClassifyRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_CREATE_CLASSIFY,adminAuth,classifyCreate.CreateClassifyController)
}
