package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/wxDispose"
	"go-server-template/pkg/apiMap"
)

func WXDisposeRouter(g *gin.RouterGroup) {
	g.POST(apiMap.POST_USER_WX_GET_OPEN_ID, wxDispose.WXGetOpenIdController)
	g.POST(apiMap.POST_USER_WX_DNCRYPT, wxDispose.WXGetOpenIdController)
}
