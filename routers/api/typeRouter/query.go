package typeRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/classifyType"
)

func QueryTypeRouter(g *gin.RouterGroup) {
	g.GET("/query-type", classifyType.QueryTypeController)
}
