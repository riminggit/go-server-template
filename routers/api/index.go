package api

import (
	"github.com/gin-gonic/gin"
	"go-server-template/routers/api/classifyRouter"
	"go-server-template/routers/api/userRouter"
	"go-server-template/routers/api/typeRouter"
	"go-server-template/routers/api/companyRouter"
	"go-server-template/routers/api/tagRouter"
	"go-server-template/routers/api/userTopicRouter"
	"go-server-template/routers/api/topicRouter"
	"go-server-template/pkg/apiMap"
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
}
