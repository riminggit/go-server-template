package userRouter

import (
	"github.com/gin-gonic/gin"
	// "go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func UserInitRouter(r *gin.RouterGroup) {
	userLogin := r.Group("/user-login")
	// userAuth.Use(JWTMiddleware.JWT())
	UserLoginRouter(userLogin)
}
