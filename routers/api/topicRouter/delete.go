package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/topic/delete"
	"go-server-template/src/topicSet/delete"
)

func DeleteTopicRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_DELETE_TOPIC, adminAuth, topicDelete.DeleteController)
	g.POST(apiMap.POST_DELETE_TOPIC_SET, adminAuth, topicSetDelete.DeleteController)
}
