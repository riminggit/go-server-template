package tagRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func TagInitRouter(r *gin.RouterGroup) {
	tag := r.Group("/tag")
	tag.Use(JWTMiddleware.JWT())
	QueryTagRouter(tag)
}
