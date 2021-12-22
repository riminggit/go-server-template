package companyRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/company/query"
	"go-server-template/pkg/apiMap"
)

func QueryCompanyRouter(g *gin.RouterGroup) {
	g.GET(apiMap.GET_QUERY_COMPANY, companyQuery.QueryCompanyController)
	g.POST(apiMap.POST_QUERY_COMPANY, companyQuery.QueryCompanyMultipleController)
}
