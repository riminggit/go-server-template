package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"
)

// 初始化模块路由
func UserInitRouter(r *gin.RouterGroup) {
	userLogin := r.Group(apiMap.USER_LOGIN_PREFIX)
	userAccount := r.Group(apiMap.USER_ACCOUNT_PREFIX)
	user := r.Group(apiMap.USER_PREFIX)



	// 添加jwt校验
	userAccount.Use(JWTMiddleware.JWT())
	user.Use(JWTMiddleware.JWT())

	UserLoginRouter(userLogin)
	WXDisposeRouter(userLogin)

	AccountManageRouter(userAccount)

	QueryUserRouter(user)
}
