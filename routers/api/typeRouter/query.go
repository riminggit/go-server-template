package typeRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/classifyType/query"
)

func QueryTypeRouter(g *gin.RouterGroup) {
	g.GET("/query-type", classifyTypeQuery.QueryTypeController)
	g.POST("/query-type", classifyTypeQuery.QueryTypeMultipleController)
}
