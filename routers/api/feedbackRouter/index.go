package feedbackRouter

import (
	"github.com/gin-gonic/gin"
	// "go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func FeedbackInitRouter(r *gin.RouterGroup) {
	tag := r.Group("/tag")
	// userAuth.Use(JWTMiddleware.JWT())
	QueryFeedbackRouter(tag)
}
