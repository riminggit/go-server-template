package companyRouter

import (
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/company/update"
	"github.com/gin-gonic/gin"
)

func UpdateCompanyRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_UPDATE_COMPANY,adminAuth,companyUpdate.UpdateController)
}
