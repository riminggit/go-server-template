package userCollectRouter

import (
	"github.com/gin-gonic/gin"
	// "go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func UserCollectRouterInitRouter(r *gin.RouterGroup) {
	userCollect := r.Group("/userCollect")
	// userAuth.Use(JWTMiddleware.JWT())
	QueryUserCollectRouter(userCollect)
}
