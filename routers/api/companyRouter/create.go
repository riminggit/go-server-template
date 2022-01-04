package companyRouter

import (
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/company/create"
	"github.com/gin-gonic/gin"
)

func CreateCompanyRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_CREATE_COMPANY,adminAuth,companyCreate.CreateController)
}
