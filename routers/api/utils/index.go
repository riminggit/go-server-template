package utils

import (
	JWTMiddleware "go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"

	"github.com/gin-gonic/gin"
)

// 初始化模块路由
func UtilsInitRouter(r *gin.RouterGroup) {
	Utils := r.Group(apiMap.UTILS_PREFIX)
	Utils.Use(JWTMiddleware.JWT())
	Color(Utils)

}
