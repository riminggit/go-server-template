package companyRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func CompanyInitRouter(r *gin.RouterGroup) {
	company := r.Group("/company")
	company.Use(JWTMiddleware.JWT())
	QueryCompanyRouter(company)
}
