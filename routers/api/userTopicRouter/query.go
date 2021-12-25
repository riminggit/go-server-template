package userTopicRouter

import (
	"go-server-template/pkg/apiMap"
	userAnswerTopicQuery "go-server-template/src/userAnswerTopic/query"
	userTopicQuery "go-server-template/src/userTopic/query"

	"github.com/gin-gonic/gin"
)

func QueryUserTopicRouter(g *gin.RouterGroup) {
	g.POST(apiMap.POST_USER_ADD_TOPIC_QUERY, userTopicQuery.QueryUserTopicController)
	g.POST(apiMap.POST_USER_ANSWER_TOPIC_QUERY, userAnswerTopicQuery.QueryUserAnswerTopicController)
}
