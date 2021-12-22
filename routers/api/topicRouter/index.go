package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"
)

// 初始化模块路由
func TopicInitRouter(r *gin.RouterGroup) {
	topic := r.Group(apiMap.TOPIC_PREFIX)
	topic.Use(JWTMiddleware.JWT())
	QueryTopicRouter(topic)
}
