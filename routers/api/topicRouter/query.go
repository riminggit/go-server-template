package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/topicSet/query"
	"go-server-template/src/topic/query"
)

func QueryTopicRouter(g *gin.RouterGroup) {
	g.POST("/query-topic-set", topicSetQuery.QueryTopicSetController)
	g.POST("/query-topic", topicQuery.QueryTopicController)
}
