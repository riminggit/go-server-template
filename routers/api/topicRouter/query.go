package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/topicSet"
)

func QueryTopicRouter(g *gin.RouterGroup) {
	g.POST("/query-topic-set", topicSet.QueryTopicSetController)
}
