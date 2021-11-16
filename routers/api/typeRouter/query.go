package typeRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/topicType"
)

func QueryTypeRouter(g *gin.RouterGroup) {
	g.GET("/query-type", topicType.QueryTypeController)
}
