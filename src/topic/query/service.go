package topicQuery

import (
	"encoding/json"
	"go-server-template/config"
	"go-server-template/model/classify"
	"go-server-template/model/company"
	"go-server-template/model/tag"
	"go-server-template/model/topic"
	"go-server-template/model/type"
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

func QueryTopicService(c *gin.Context, params QueryTopicParams) *queryTopicReturn {
	res := &queryTopicReturn{}
	var queryInfo []topicModel.Topic
	var RGetData TopicReturnData

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &RGetData)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.Data = RGetData
		return res
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)
	if len(params.Id) > 0 {
		queryFun = queryFun.Where("id IN (?)", params.Id)
	}

	if params.Title != "" {
		queryFun = queryFun.Where("title = ?", params.Title)
	}

	if len(params.QuestionType) > 0 {
		queryFun = queryFun.Where("question_type IN (?)", params.QuestionType)
	}

	if len(params.Degree) > 0 {
		queryFun = queryFun.Where("degree IN (?)", params.Degree)
	}

	if len(params.Level) > 0 {
		queryFun = queryFun.Where("level IN (?)", params.Level)
	}

	if params.IsBaseTopic != "" {
		queryFun = queryFun.Where("is_base_topic = ?", params.IsBaseTopic)
	}

	if params.IsImportantTopic != "" {
		queryFun = queryFun.Where("is_important_topic = ?", params.IsImportantTopic)
	}

	if len(params.ComeFrom) > 0 {
		queryFun = queryFun.Where("come_from IN (?)", params.ComeFrom)
	}

	if len(params.CreateAt) > 0 {
		queryFun = queryFun.Where("create_at between ? and ?", params.CreateAt[0], params.CreateAt[1])
	}

	if len(params.DeleteAt) > 0 {
		queryFun = queryFun.Where("delete_at between ? and ?", params.DeleteAt[0], params.DeleteAt[1])
	}

	if len(params.UpdateAt) > 0 {
		queryFun = queryFun.Where("update_at between ? and ?", params.UpdateAt[0], params.UpdateAt[1])
	}

	queryFun = queryFun.Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize)
	queryFun.Model(&topicModel.Topic{}).Find(&queryInfo).Count(&res.Data.PagingArgument.Total)

	res.Data.PagingArgument.PageNum = params.PageNum
	res.Data.PagingArgument.PageSize = params.PageSize
	res.Data.Data = queryInfo

	RSetData := res.Data
	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}

func QueryTopicRelationService(c *gin.Context, params QueryTopicParams) *queryTopicReturnRelation {
	res := &queryTopicReturnRelation{}

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_RELACTION)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	var topicInfo TopicReturnDataRelation

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &topicInfo)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.Data = topicInfo
		return res
	}

	result := QueryTopicService(c, params)

	for _, item := range result.Data.Data {
		var classifyInfo []classifyModel.Classify
		var companyInfo []companyModel.Company
		var tagInfo []tagModel.Tag
		var typeInfo []typeModel.Type
		itemReturnData := TopicData{
			item,
			classifyInfo,
			companyInfo,
			tagInfo,
			typeInfo,
		}

		for _, judgeItem := range params.Relation {
			if judgeItem == "classify" {
				// classify
				CParams := midTopicClassifyQuery.QueryTopicClassifyMidParams{
					TopicId: strconv.Itoa(item.ID),
				}
				queryClassifyMidData := midTopicClassifyQuery.QueryTopicClassifyMid(c, CParams).Data
				if len(queryClassifyMidData) > 0 {
					for _, classifyItem := range queryClassifyMidData {
						classifyParams := classifyQuery.QueryClassifyParams{
							Id:    strconv.Itoa(classifyItem.ClassifyId),
							IsUse: "1",
						}
						queryClassifyData := classifyQuery.QueryClassifyService(c, classifyParams).Data
						itemReturnData.ClassifyInfo = append(itemReturnData.ClassifyInfo, queryClassifyData...)
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
						itemReturnData.CompanyInfo = append(itemReturnData.CompanyInfo, queryCompanyData...)
					}
				}
			}

			if judgeItem == "tag" {
				// tag
				TagParams := midTopicTagQuery.QueryTopicTagMidParams{
					TopicId: strconv.Itoa(item.ID),
				}
				queryTagMidData := midTopicTagQuery.QueryTopicTagMid(c, TagParams).Data
				if len(queryTagMidData) > 0 {
					for _, tagItem := range queryTagMidData {
						tagParams := tagQuery.QueryTagParams{
							Id:    strconv.Itoa(tagItem.TagId),
							IsUse: "1",
						}
						queryTagData := tagQuery.QueryTagService(c, tagParams).Data
						itemReturnData.TagInfo = append(itemReturnData.TagInfo, queryTagData...)
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
						itemReturnData.TypeInfo = append(itemReturnData.TypeInfo, queryTypeData...)
					}
				}
			}

		}

		topicInfo.Data = append(topicInfo.Data, itemReturnData)
	}

	topicInfo.PagingArgument.PageNum = result.Data.PagingArgument.PageNum
	topicInfo.PagingArgument.PageSize = result.Data.PagingArgument.PageSize
	topicInfo.PagingArgument.Total = result.Data.PagingArgument.Total
	res.Data = topicInfo

	RSetData := topicInfo
	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}



func QueryTopicSimpleService(c *gin.Context, params QueryTopicSimpleParams) *TopicSimpleReturn {
	res := &TopicSimpleReturn{}
	res.Code = e.SUCCESS
	var queryInfo []topicModel.Topic

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_SIMPLE)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &queryInfo)
		if err != nil {
			logging.Debug(err)
		}
		res.Data = queryInfo
		return res
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", 1)
	if len(params.Id) > 0 {
		queryFun = queryFun.Where("id IN (?)", params.Id)
	}

	if params.Title != "" {
		queryFun = queryFun.Where("title = ?", params.Title)
	}

	resp := queryFun.Model(&topicModel.Topic{}).Find(&queryInfo)
	if resp.Error != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)
	res.Data = queryInfo

	return res
}



func QueryTopicSimpleSingleService(c *gin.Context, params QueryTopicSimpleSingleParams) *TopicSimpleSingleReturn {
	res := &TopicSimpleSingleReturn{}
	res.Code = e.SUCCESS
	var queryInfo topicModel.Topic

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME
	redisParamsJson, _ := json.Marshal(params)
	interfaceName := apiMap.GetRedisPrefixName(apiMap.POST_QUERY_TOPIC_SIMPLE_SINGLE)
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &queryInfo)
		if err != nil {
			logging.Debug(err)
		}
		res.Data = queryInfo
		return res
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", 1)
	if params.Id > 0 {
		queryFun = queryFun.Where("id = ?", params.Id)
	}

	if params.Title != "" {
		queryFun = queryFun.Where("title = ?", params.Title)
	}

	resp := queryFun.Model(&topicModel.Topic{}).Find(&queryInfo)
	if resp.Error != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}

	Redis.SetValue(queryRedisParams, queryInfo, dataRxpirationTime)
	res.Data = queryInfo

	return res
}