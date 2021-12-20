package classifyRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/classify/query"
)

func QueryClassifyRouter(g *gin.RouterGroup) {
	g.GET("/query-classify", classifyQuery.QueryClassifyController)
	g.GET("/query-classify-type", classifyQuery.QueryClassifyAndTypeController)
	g.POST("/query-classify", classifyQuery.QueryClassifyMultipleController)
	g.POST("/query-classify-type", classifyQuery.QueryClassifyAndTypeMultipleController)
}
