package classifyRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/classify"
)

func QueryClassifyRouter(g *gin.RouterGroup) {
	g.POST("/query-classify", classify.QueryClassifyController)
}
