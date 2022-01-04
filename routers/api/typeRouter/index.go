package typeRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"
)

// 初始化模块路由
func TypeInitRouter(r *gin.RouterGroup) {
	typeRouter := r.Group(apiMap.TYPE_PREFIX)
	typeRouter.Use(JWTMiddleware.JWT())
	QueryTypeRouter(typeRouter)
	CreateTypeRouter(typeRouter)
	DeleteTypeRouter(typeRouter)
	UpdateTypeRouter(typeRouter)
}
