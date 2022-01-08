package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/topic/create"
	"go-server-template/src/topicSet/create"
)

func CreateTopicRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_CREATE_TOPIC, adminAuth, topicCreate.CreateController)
	g.POST(apiMap.POST_CREATE_TOPIC_SET, adminAuth, topicSetCreate.CreateController)
}
