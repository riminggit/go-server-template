package userTopicRouter

import (
	"go-server-template/pkg/apiMap"
	userTopicQuery "go-server-template/src/userTopic/query"

	"github.com/gin-gonic/gin"
)

func UserAddTopic(g *gin.RouterGroup) {
	// 查询用户自己新增的题目
	g.POST(apiMap.POST_USER_ADD_TOPIC_QUERY, userTopicQuery.QueryUserTopicController)
}
