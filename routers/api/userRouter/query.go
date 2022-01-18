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
	// 管理员查询用户经验
	g.POST(apiMap.POST_QUERY_USER_EXPERIENCE, adminAuth, userExperience.QueryExperienceController)
	// 管理员查询用户信息
	g.POST(apiMap.POST_ADMIN_QUERY_USER_INFO, adminAuth, userQuery.QueryUserDataController)

	// 用户查询经验
	g.GET(apiMap.POST_QUERY_USER_EXPERIENCE_FROM_USER, userExperience.UserQueryExperienceController)

}
