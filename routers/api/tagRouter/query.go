package tagRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/tag"
)

func QueryTagRouter(g *gin.RouterGroup) {
	g.GET("/query-tag", tag.QueryTagController)
}
