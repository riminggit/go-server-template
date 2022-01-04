package typeRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/classifyType/update"
	"go-server-template/pkg/apiMap"
	"go-server-template/middleware/global/auth"
)

func UpdateTypeRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_UPDATE_TYPE,adminAuth,classifyTypeUpdate.UpdateController)
}
