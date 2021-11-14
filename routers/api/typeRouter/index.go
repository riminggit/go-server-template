package typeRouter

import (
	"github.com/gin-gonic/gin"
	// "go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func TypeInitRouter(r *gin.RouterGroup) {
	typeRouter := r.Group("/type")
	// userAuth.Use(JWTMiddleware.JWT())
	QueryTypeRouter(typeRouter)
}
