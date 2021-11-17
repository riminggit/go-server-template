package companyRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/company"
)

func QueryCompanyRouter(g *gin.RouterGroup) {
	g.GET("/query-company", company.QueryCompanyController)
}
