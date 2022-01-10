package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/user/query"
	"go-server-template/src/userExperience"
)

func QueryUserRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_ADMIN_QUERY_USER_INFO, userQuery.QueryUserDataController)
	g.POST(apiMap.POST_QUERY_USER_EXPERIENCE, adminAuth, userExperience.QueryExperienceController)

}
