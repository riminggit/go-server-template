package userRouter

import (
	"go-server-template/src/wxDispose"

	"github.com/gin-gonic/gin"
)

func WXDisposeRouter(g *gin.RouterGroup) {
	g.POST("/wxapp-get-openid", wxDispose.WXGetOpenIdController)
	g.POST("/wxapp-dncrypt", wxDispose.WXGetOpenIdController)

}
