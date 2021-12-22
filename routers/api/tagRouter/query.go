package tagRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/tag/query"
	"go-server-template/pkg/apiMap"
)

func QueryTagRouter(g *gin.RouterGroup) {
	g.GET(apiMap.GET_QUERY_TAG, tagQuery.QueryTagController)
	g.POST(apiMap.POST_QUERY_TAG, tagQuery.QueryTagMultipleController)
}
