package userTopicRouter

import (
	"go-server-template/pkg/apiMap"
	userAnswerTopicCreate "go-server-template/src/userAnswerTopic/create"
	userAnswerTopicQuery "go-server-template/src/userAnswerTopic/query"
	userAnswerTopicUpdate "go-server-template/src/userAnswerTopic/update"

	"github.com/gin-gonic/gin"
)

func UserAnswerRecord(g *gin.RouterGroup) {
	// 查询用户答题记录
	g.POST(apiMap.POST_USER_ANSWER_TOPIC_QUERY, userAnswerTopicQuery.Controller)
	// 新增用户答题记录
	g.POST(apiMap.POST_USER_ANSWER_TOPIC_CREATE, userAnswerTopicCreate.Controller)
	// 修改用户答题记录
	g.POST(apiMap.POST_USER_ANSWER_TOPIC_UPDATE, userAnswerTopicUpdate.Controller)
}
