package api

import (
	"go-server-template/pkg/apiMap"
	"go-server-template/routers/api/classifyRouter"
	"go-server-template/routers/api/companyRouter"
	"go-server-template/routers/api/feedbackRouter"
	"go-server-template/routers/api/tagRouter"
	"go-server-template/routers/api/topicRouter"
	"go-server-template/routers/api/typeRouter"
	"go-server-template/routers/api/userRouter"
	"go-server-template/routers/api/userTopicRouter"
	"go-server-template/routers/api/utils"

	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {
	api := r.Group(apiMap.API_PREFIX)
	userRouter.UserInitRouter(api)
	classifyRouter.ClassifyInitRouter(api)
	typeRouter.TypeInitRouter(api)
	companyRouter.CompanyInitRouter(api)
	tagRouter.TagInitRouter(api)
	userTopicRouter.UserTopicInitRouter(api)
	topicRouter.TopicInitRouter(api)
	feedbackRouter.FeedbackInitRouter(api)
	utils.UtilsInitRouter(api)
}
