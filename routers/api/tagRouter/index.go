package tagRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"
)

// 初始化模块路由
func TagInitRouter(r *gin.RouterGroup) {
	tag := r.Group(apiMap.TAG_PREFIX)
	tag.Use(JWTMiddleware.JWT())
	QueryTagRouter(tag)
}
