package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/topicSet/query"
	"go-server-template/src/topic/query"
	"go-server-template/src/midTopicClassify/query"
	"go-server-template/src/midTopicCompany/query"
	"go-server-template/src/midTopicTag/query"
	"go-server-template/src/midTopicType/query"
)

func QueryTopicRouter(g *gin.RouterGroup) {
	g.POST("/query-topic-set", topicSetQuery.QueryTopicSetController)
	g.POST("/query-topic", topicQuery.QueryTopicController)
	g.POST("/query-topic/classify/mid", midTopicClassifyQuery.QueryTopicClassifyMidController)
	g.POST("/query-topic/company/mid", midTopicCompanyQuery.QueryTopicCompanyMidController)
	g.POST("query-topic/tag/mid", midTopicTagQuery.QueryTopicTagMidController)
	g.POST("query-topic/type/mid", midTopicTypeQuery.QueryTopicTypeMidController)
	
}
