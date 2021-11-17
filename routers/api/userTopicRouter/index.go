package userTopicRouter

import (
	"github.com/gin-gonic/gin"
	// "go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func UserTopicInitRouter(r *gin.RouterGroup) {
	userTopic := r.Group("/user-topic")
	// userAuth.Use(JWTMiddleware.JWT())
	QueryUserTopicRouter(userTopic)
}
