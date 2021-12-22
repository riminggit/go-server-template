package userTopicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"
)

// 初始化模块路由
func UserTopicInitRouter(r *gin.RouterGroup) {
	userTopic := r.Group(apiMap.USER_TOPIC_PREFIX)
	userTopic.Use(JWTMiddleware.JWT())
	QueryUserTopicRouter(userTopic)
}
