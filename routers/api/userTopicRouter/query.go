package userTopicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/userTopic"
)

func QueryUserTopicRouter(g *gin.RouterGroup) {
	g.POST("/query-user-topic", userTopic.QueryUserTopicController)
}
