package userTopicRouter

import (
	"go-server-template/pkg/apiMap"
	"go-server-template/src/userTopicRead"

	"github.com/gin-gonic/gin"
)

func UserTopicReadRecord(g *gin.RouterGroup) {
	// 新增用户读题记录
	// 查询题目五秒后请求
	g.POST(apiMap.POST_USER_TOPIC_READ_RECORD_CREATE, userTopicRead.CreateController)

}
