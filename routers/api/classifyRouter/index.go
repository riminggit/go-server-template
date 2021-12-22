package classifyRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"
)

// 初始化模块路由
func ClassifyInitRouter(r *gin.RouterGroup) {
	clssify := r.Group(apiMap.CLASSIFY_PREFIX)
	clssify.Use(JWTMiddleware.JWT())
	QueryClassifyRouter(clssify)
}
