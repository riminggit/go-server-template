package classifyRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/classify/query"
	"go-server-template/pkg/apiMap"
)

func QueryClassifyRouter(g *gin.RouterGroup) {
	g.GET(apiMap.GET_QUERY_CLASSIFY, classifyQuery.QueryClassifyController)
	g.GET(apiMap.GET_QUERY_CLASSIFY_TYPE, classifyQuery.QueryClassifyAndTypeController)
	g.POST(apiMap.POST_QUERY_CLASSIFY, classifyQuery.QueryClassifyMultipleController)
	g.POST(apiMap.POST_QUERY_CLASSIFY_TYPE, classifyQuery.QueryClassifyAndTypeMultipleController)
}
