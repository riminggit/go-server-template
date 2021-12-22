package topicRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/topicSet/query"
	"go-server-template/src/topic/query"
	"go-server-template/src/midTopicClassify/query"
	"go-server-template/src/midTopicCompany/query"
	"go-server-template/src/midTopicTag/query"
	"go-server-template/src/midTopicType/query"
	"go-server-template/pkg/apiMap"
)

func QueryTopicRouter(g *gin.RouterGroup) {
	g.POST(apiMap.POST_QUERY_TOPIC_SET, topicSetQuery.QueryTopicSetController)
	g.POST(apiMap.POST_QUERY_TOPIC, topicQuery.QueryTopicController)
	g.POST(apiMap.POST_QUERY_TOPIC_CLASSIFY_MID, midTopicClassifyQuery.QueryTopicClassifyMidController)
	g.POST(apiMap.POST_QUERY_TOPIC_COMPANY_MID, midTopicCompanyQuery.QueryTopicCompanyMidController)
	g.POST(apiMap.POST_QUERY_TOPIC_TAG_MID, midTopicTagQuery.QueryTopicTagMidController)
	g.POST(apiMap.POST_QUERY_TOPIC_TYPE_MID, midTopicTypeQuery.QueryTopicTypeMidController)
	
}
