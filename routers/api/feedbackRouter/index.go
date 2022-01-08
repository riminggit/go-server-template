package feedbackRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"
)

// 初始化模块路由
func FeedbackInitRouter(r *gin.RouterGroup) {
	feedback := r.Group(apiMap.FEEDBACK_PREFIX)
	feedback.Use(JWTMiddleware.JWT())
	userFeedbackRouter(feedback)
}
