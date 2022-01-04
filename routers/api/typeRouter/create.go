package typeRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/classifyType/create"
)

func CreateTypeRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_CREATE_TYPE, adminAuth, classifyTypeCreate.CreateTypeController)
}
