package utils

import (
	authMiddleware "go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/utils"

	"github.com/gin-gonic/gin"
)

func Color(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	// 新增颜色
	g.POST(apiMap.POST_COLOR_ADD, adminAuth, utils.AddColor)
	g.GET(apiMap.POST_COLOR_QUERY, utils.QueryColor)
}
