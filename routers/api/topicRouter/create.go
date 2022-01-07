package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/topic/create"
)

func CreateTopicRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_CREATE_TOPIC, adminAuth, topicCreate.CreateController)
}
