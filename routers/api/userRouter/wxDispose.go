package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/wxDispose"
)

func WXDisposeRouter(g *gin.RouterGroup) {
	g.POST("/wxapp-get-openid", wxDispose.WXGetOpenIdController)
	g.POST("/wxapp-dncrypt", wxDispose.WXGetOpenIdController)

}
