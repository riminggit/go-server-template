package userTopicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/userTopic/query"
	"go-server-template/src/userAnswerTopic/query"
	"go-server-template/pkg/apiMap"
)

func QueryUserTopicRouter(g *gin.RouterGroup) {
	g.POST(apiMap.POST_USER_TOPIC_QUERY, userTopicQuery.QueryUserTopicController)
	g.POST(apiMap.POST_USER_ANSWER_TOPIC_QUERY, userAnswerTopicQuery.QueryUserAnswerTopicController)
}
   