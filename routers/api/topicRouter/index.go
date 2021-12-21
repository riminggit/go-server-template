package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func TopicInitRouter(r *gin.RouterGroup) {
	topic := r.Group("/topic")
	topic.Use(JWTMiddleware.JWT())
	QueryTopicRouter(topic)
}
