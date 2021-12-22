package companyRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
	"go-server-template/pkg/apiMap"
)

// 初始化模块路由
func CompanyInitRouter(r *gin.RouterGroup) {
	company := r.Group(apiMap.COMPANY_PREFIX)
	company.Use(JWTMiddleware.JWT())
	QueryCompanyRouter(company)
}
