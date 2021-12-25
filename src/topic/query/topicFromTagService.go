package topicQuery

import (
	"encoding/json"
	"go-server-template/config"
	"go-server-template/model/topic"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	"go-server-template/src/classify/query"
	"go-server-template/src/classifyType/query"
	"go-server-template/src/company/query"
	"go-server-template/src/midTopicClassify/query"
	"go-server-template/src/midTopicCompany/query"
	"go-server-template/src/midTopicTag/query"
	"go-server-template/src/midTopicType/query"
	"go-server-template/src/tag/query"
	"strconv"
	"github.com/gin-gonic/gin"
)

func QueryTopicIdFromTag(c *gin.Context, params *queryTopicFromTagParams) *QueryTopicIdReturn {
	res := &QueryTopicIdReturn{}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_FROM_TAG_GET_TOPIC_ID)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	var topicIdList []string

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &topicIdList)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.TopicIdList = topicIdList
		return res
	}

	for _, item := range params.TagId {
		TParams := midTopicTagQuery.QueryTopicTagMidParams{
			TagId: item,
		}
		queryMidData := midTopicTagQuery.QueryTopicTagMid(c, TParams)
		if queryMidData.Code == e.SUCCESS && len(queryMidData.Data) > 0 {
			for _, item2 := range queryMidData.Data {
				topicIdList = append(topicIdList, strconv.Itoa(item2.TopicId))
			}
		} else {
			res.Code = e.ERROR
			return res
		}
	}

	res.TopicIdList = topicIdList
	Redis.SetValue(queryRedisParams, topicIdList, dataRxpirationTime)
	res.Code = e.SUCCESS
	return res
}

func QueryTopicFromTagService(c *gin.Context, params queryTopicFromTagParams) *queryTopicNoPadingReturn {
	res := &queryTopicNoPadingReturn{}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_FROM_TAG)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	var topicInfo []topicModel.Topic

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &topicInfo)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.Data = topicInfo
		return res
	}

	result := QueryTopicIdFromTag(c, &params)
	if result.Code == e.SUCCESS {
		queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)
		queryFun = queryFun.Where("id IN (?)", result.TopicIdList)
		queryFun.Model(&topicModel.Topic{}).Find(&topicInfo)
		res.Data = topicInfo
		Redis.SetValue(queryRedisParams, topicInfo, dataRxpirationTime)
		res.Code = e.SUCCESS
		return res
	}

	res.Code = e.NO_DATA_EXISTS
	return res
}

func QueryTopicFromTagRelationService(c *gin.Context, params queryTopicFromTagParams) *queryTopicNoPadingRelationReturn {
	res := &queryTopicNoPadingRelationReturn{}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_FROM_TAG_RELATION)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	var topicInfo []TopicData
	var topicInfoReturn []TopicData

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &topicInfo)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.Data = topicInfo
		return res
	}

	result := QueryTopicIdFromTag(c, &params)
	if result.Code == e.SUCCESS {
		queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)
		queryFun = queryFun.Where("id IN (?)", result.TopicIdList)
		queryFun.Model(&topicModel.Topic{}).Find(&topicInfo)

		for _, item := range topicInfo {
			for _, judgeItem := range params.Relation {
				if judgeItem == "tag" {
					for _, item2 := range params.TagId {
						thisParams := tagQuery.QueryTagParams{
							Id:    item2,
							IsUse: "1",
						}
						queryData := tagQuery.QueryTagService(c, thisParams).Data
						item.TagInfo = append(item.TagInfo, queryData...)
					}
				}

				if judgeItem == "classify" {
					thisParams := midTopicClassifyQuery.QueryTopicClassifyMidParams{
						TopicId: strconv.Itoa(item.ID),
					}
					queryClassifyMidData := midTopicClassifyQuery.QueryTopicClassifyMid(c, thisParams).Data
					if len(queryClassifyMidData) > 0 {
						for _, classifyItem := range queryClassifyMidData {
							classifyParams := classifyQuery.QueryClassifyParams{
								Id:    strconv.Itoa(classifyItem.ClassifyId),
								IsUse: "1",
							}
							queryClassifyData := classifyQuery.QueryClassifyService(c, classifyParams).Data
							item.ClassifyInfo = append(item.ClassifyInfo, queryClassifyData...)
						}
					}
				}

				if judgeItem == "company" {
					// company
					CompanyParams := midTopicCompanyQuery.QueryTopicCompanyMidParams{
						TopicId: strconv.Itoa(item.ID),
					}
					queryCompanyMidData := midTopicCompanyQuery.QueryTopicCompanyMid(c, CompanyParams).Data
					if len(queryCompanyMidData) > 0 {
						for _, companyItem := range queryCompanyMidData {
							companyParams := companyQuery.QueryCompanyParams{
								Id:    strconv.Itoa(companyItem.CompanyId),
								IsUse: "1",
							}
							queryCompanyData := companyQuery.QueryCompanyService(c, companyParams).Data
							item.CompanyInfo = append(item.CompanyInfo, queryCompanyData...)
						}
					}
				}

				if judgeItem == "type" {
					// type
					TypeParams := midTopicTypeQuery.QueryTopicTypeMidParams{
						TopicId: strconv.Itoa(item.ID),
					}
					queryTypeMidData := midTopicTypeQuery.QueryTopicTypeMid(c, TypeParams).Data
					if len(queryTypeMidData) > 0 {
						for _, typeItem := range queryTypeMidData {
							typeParams := classifyTypeQuery.QueryTypeParams{
								Id:    strconv.Itoa(typeItem.TypeId),
								IsUse: "1",
							}
							queryTypeData := classifyTypeQuery.QueryTypeService(c, typeParams).Data
							item.TypeInfo = append(item.TypeInfo, queryTypeData...)
						}
					}
				}
			}
			topicInfoReturn = append(topicInfoReturn, item)
		}
		res.Data = topicInfoReturn
		Redis.SetValue(queryRedisParams, topicInfoReturn, dataRxpirationTime)
		res.Code = e.SUCCESS
		return res
	}

	res.Code = e.NO_DATA_EXISTS
	return res
}
