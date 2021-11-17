package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func UserInitRouter(r *gin.RouterGroup) {
	userLogin := r.Group("/user-login")
	userAccount := r.Group("/user/account-manage")
	user := r.Group("/user")



	// 添加jwt校验
	userAccount.Use(JWTMiddleware.JWT())
	user.Use(JWTMiddleware.JWT())

	UserLoginRouter(userLogin)
	WXDisposeRouter(userLogin)

	AccountManageRouter(userAccount)

	QueryUserRouter(user)
}
