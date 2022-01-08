package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/topic/update"
	"go-server-template/src/topicSet/update"
)

func UpdateTopicRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_UPDATE_TOPIC, adminAuth, topicUpdate.UpdateController)
	g.POST(apiMap.POST_UPDATE_TOPIC_SET, adminAuth, topicSetUpdate.UpdateController)
}
