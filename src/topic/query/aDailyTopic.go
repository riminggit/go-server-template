package topicQuery

import (
	"encoding/json"
	projectConfig "go-server-template/config"
	classifyModel "go-server-template/model/classify"
	companyModel "go-server-template/model/company"
	tagModel "go-server-template/model/tag"
	topicModel "go-server-template/model/topic"
	typeModel "go-server-template/model/type"
	"go-server-template/pkg/apiMap"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/pkg/level"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	util "go-server-template/pkg/utils"
	classifyQuery "go-server-template/src/classify/query"
	classifyTypeQuery "go-server-template/src/classifyType/query"
	companyQuery "go-server-template/src/company/query"
	midTopicClassifyQuery "go-server-template/src/midTopicClassify/query"
	midTopicCompanyQuery "go-server-template/src/midTopicCompany/query"
	midTopicTagQuery "go-server-template/src/midTopicTag/query"
	midTopicTypeQuery "go-server-template/src/midTopicType/query"
	tagQuery "go-server-template/src/tag/query"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

func GetADailyTopic(c *gin.Context) *QueryADailyTopicReturn {
	res := &QueryADailyTopicReturn{
		Code: e.SUCCESS,
	}

	resData := &TopicData{}
	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME / 2

	interfaceName := apiMap.GetRedisPrefixName(apiMap.GET_A_DAILY_TOPIC)
	queryRedisParams := interfaceName + time.Now().Format("2006-01-02")

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &resData)
		if err != nil {
			logging.Debug(err)
		}

		res.Data = *resData
		return res
	}

	res.Data = *GetADailyTopicRand(c)

	RSetData := res.Data
	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	return res
}

func GetADailyTopicRand(c *gin.Context) *TopicData {
	result := &TopicData{}
	topicData := &topicModel.Topic{}
	queryCount := "1"
	tableName := `topic`
	levelHelper := 3

	num := util.RandNum()
	if num >= 0 && num <= 10 {
		levelHelper = level.PRIMARY
	} else if num <= 30 {
		levelHelper = level.MIDDLE_RANK
	} else if num <= 75 {
		levelHelper = level.ADVANCED
	} else if num <= 100 {
		levelHelper = level.SENIOR
	}

	sqlHelper := "SELECT * FROM " + tableName + " WHERE level = " + strconv.Itoa(levelHelper) + " AND is_use = 1 ORDER BY RAND() LIMIT " + queryCount
	DB.DBLivingExample.Raw(sqlHelper).Scan(&topicData)

	var classifyInfo []classifyModel.Classify
	var companyInfo []companyModel.Company
	var tagInfo []tagModel.Tag
	var typeInfo []typeModel.Type
	returnData := TopicData{
		*topicData,
		classifyInfo,
		companyInfo,
		tagInfo,
		typeInfo,
	}

	// classify
	CParams := midTopicClassifyQuery.QueryTopicClassifyMidParams{
		TopicId: strconv.Itoa(topicData.ID),
	}
	queryClassifyMidData := midTopicClassifyQuery.QueryTopicClassifyMid(c, CParams).Data
	if len(queryClassifyMidData) > 0 {
		for _, classifyItem := range queryClassifyMidData {
			classifyParams := classifyQuery.QueryClassifyParams{
				Id:    strconv.Itoa(classifyItem.ClassifyId),
				IsUse: "1",
			}
			queryClassifyData := classifyQuery.QueryClassifyService(c, classifyParams).Data
			returnData.ClassifyInfo = append(returnData.ClassifyInfo, queryClassifyData...)
		}
	}

	// company
	CompanyParams := midTopicCompanyQuery.QueryTopicCompanyMidParams{
		TopicId: strconv.Itoa(topicData.ID),
	}
	queryCompanyMidData := midTopicCompanyQuery.QueryTopicCompanyMid(c, CompanyParams).Data
	if len(queryCompanyMidData) > 0 {
		for _, companyItem := range queryCompanyMidData {
			companyParams := companyQuery.QueryCompanyParams{
				Id:    strconv.Itoa(companyItem.CompanyId),
				IsUse: "1",
			}
			queryCompanyData := companyQuery.QueryCompanyService(c, companyParams).Data
			returnData.CompanyInfo = append(returnData.CompanyInfo, queryCompanyData...)
		}
	}

	// tag
	TagParams := midTopicTagQuery.QueryTopicTagMidParams{
		TopicId: strconv.Itoa(topicData.ID),
	}
	queryTagMidData := midTopicTagQuery.QueryTopicTagMid(c, TagParams).Data
	if len(queryTagMidData) > 0 {
		for _, tagItem := range queryTagMidData {
			tagParams := tagQuery.QueryTagParams{
				Id:    strconv.Itoa(tagItem.TagId),
				IsUse: "1",
			}
			queryTagData := tagQuery.QueryTagService(c, tagParams).Data
			returnData.TagInfo = append(returnData.TagInfo, queryTagData...)
		}
	}

	// type
	TypeParams := midTopicTypeQuery.QueryTopicTypeMidParams{
		TopicId: strconv.Itoa(topicData.ID),
	}
	queryTypeMidData := midTopicTypeQuery.QueryTopicTypeMid(c, TypeParams).Data
	if len(queryTypeMidData) > 0 {
		for _, typeItem := range queryTypeMidData {
			typeParams := classifyTypeQuery.QueryTypeParams{
				Id:    strconv.Itoa(typeItem.TypeId),
				IsUse: "1",
			}
			queryTypeData := classifyTypeQuery.QueryTypeService(c, typeParams).Data
			returnData.TypeInfo = append(returnData.TypeInfo, queryTypeData...)
		}
	}

	result = &returnData

	if result.ID == 0 {
		result = GetADailyTopicRand(c)
	}

	return result
}
