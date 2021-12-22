package userCollectRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"
)

// 初始化模块路由
func UserCollectRouterInitRouter(r *gin.RouterGroup) {
	userCollect := r.Group(apiMap.USER_COLLECT_PREFIX)
	userCollect.Use(JWTMiddleware.JWT())
	QueryUserCollectRouter(userCollect)
}
