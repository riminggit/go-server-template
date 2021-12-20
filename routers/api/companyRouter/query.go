package companyRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/company/query"
)

func QueryCompanyRouter(g *gin.RouterGroup) {
	g.GET("/query-company", companyQuery.QueryCompanyController)
	g.POST("/query-company", companyQuery.QueryCompanyMultipleController)
}
