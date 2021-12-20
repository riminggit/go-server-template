package tagRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/tag/query"
)

func QueryTagRouter(g *gin.RouterGroup) {
	g.GET("/query-tag", tagQuery.QueryTagController)
	g.POST("/query-tag", tagQuery.QueryTagMultipleController)
}
