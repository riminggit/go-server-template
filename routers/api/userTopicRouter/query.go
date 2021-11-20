package userTopicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/userTopic"
	"go-server-template/src/userAnswerTopic"
)

func QueryUserTopicRouter(g *gin.RouterGroup) {
	g.POST("/query-user-topic", userTopic.QueryUserTopicController)
	g.POST("/query-user-answer-topic", userAnswerTopic.QueryUserAnswerTopicController)
}
