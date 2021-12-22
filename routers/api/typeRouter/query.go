package typeRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/classifyType/query"
	"go-server-template/pkg/apiMap"
)

func QueryTypeRouter(g *gin.RouterGroup) {
	g.GET(apiMap.GET_QUERY_TYPE, classifyTypeQuery.QueryTypeController)
	g.POST(apiMap.POST_QUERY_TYPE, classifyTypeQuery.QueryTypeMultipleController)
}
