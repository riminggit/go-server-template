package userTopicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/userTopic/query"
	"go-server-template/src/userAnswerTopic/query"
)

func QueryUserTopicRouter(g *gin.RouterGroup) {
	g.POST("/query-user-topic", userTopicQuery.QueryUserTopicController)
	g.POST("/query-user-answer-topic", userAnswerTopicQuery.QueryUserAnswerTopicController)
}
   