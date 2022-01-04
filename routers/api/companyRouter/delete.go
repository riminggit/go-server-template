package companyRouter

import (
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/company/delete"
	"github.com/gin-gonic/gin"
)

func DeleteCompanyRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_DELETE_COMPANY,adminAuth,companyDelete.DeleteController)
}
