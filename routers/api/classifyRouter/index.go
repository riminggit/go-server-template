package classifyRouter

import (
	"github.com/gin-gonic/gin"
	// "go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func ClassifyInitRouter(r *gin.RouterGroup) {
	clssify := r.Group("/classify")
	// userAuth.Use(JWTMiddleware.JWT())
	QueryClassifyRouter(clssify)
}
