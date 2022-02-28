package topicCreate

import (
	topicModel "go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/pkg/snowflake"
	util "go-server-template/pkg/utils"
	midTopicClassifyCreate "go-server-template/src/midTopicClassify/create"
	midTopicCompanyCreate "go-server-template/src/midTopicCompany/create"
	midTopicTagCreate "go-server-template/src/midTopicTag/create"
	midTopicTypeCreate "go-server-template/src/midTopicType/create"

	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context, params CreateParams) *CreateReturn {
	res := &CreateReturn{}
	res.Code = e.SUCCESS

	tx := DB.DBLivingExample.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			res.Code = e.CREATE_DATA_FALE
			res.Data = append(res.Data, "题目新增失败")
		}
	}()

	createTime := util.GetNowTimeUnix()

	createData := &topicModel.Topic{
		ID:               snowflake.GenerateID(1),
		Title:            params.Title,
		QuestionType:     params.QuestionType,
		Analysis:         params.Analysis,
		SelectAnalysis:   params.SelectAnalysis,
		Degree:           params.Degree,
		Level:            params.Level,
		IsBaseTopic:      params.IsBaseTopic,
		IsImportantTopic: params.IsImportantTopic,
		ComeFrom:         params.ComeFrom,
		CreateAt:         createTime,
		IsUse:            1,
	}

	err := tx.Model(&topicModel.Topic{}).Create(createData).Error
	if err != nil {
		tx.Rollback()
		res.Code = e.CREATE_DATA_FALE
	}

	var createTopicClassify []topicModel.TopicClassify
	for _, classifyItem := range params.ClassifyId {
		topicClassify := topicModel.TopicClassify{
			TopicId:    createData.ID,
			ClassifyId: classifyItem,
			CreateAt:   createTime,
		}
		createTopicClassify = append(createTopicClassify, topicClassify)
	}
	TCErr := midTopicClassifyCreate.CreateMultipleService(createTopicClassify, tx)
	if TCErr != nil {
		tx.Rollback()
		res.Code = e.CREATE_DATA_FALE
	}

	var createTopicCompany []topicModel.TopicCompany
	for _, companyItem := range params.CompanyId {
		topicCompany := topicModel.TopicCompany{
			TopicId:   createData.ID,
			CompanyId: companyItem,
			CreateAt:  createTime,
		}
		createTopicCompany = append(createTopicCompany, topicCompany)
	}
	TCompanyErr := midTopicCompanyCreate.CreateMultipleService(createTopicCompany, tx)
	if TCompanyErr != nil {
		tx.Rollback()
		res.Code = e.CREATE_DATA_FALE
	}

	var createTopicTag []topicModel.TopicTag
	for _, tagItem := range params.TagId {
		topicTag := topicModel.TopicTag{
			TopicId:  createData.ID,
			TagId:    tagItem,
			CreateAt: createTime,
		}
		createTopicTag = append(createTopicTag, topicTag)
	}

	TTagErr := midTopicTagCreate.CreateMultipleService(createTopicTag, tx)
	if TTagErr != nil {
		tx.Rollback()
		res.Code = e.CREATE_DATA_FALE
	}

	var createTopicType []topicModel.TopicType
	for _, typeItem := range params.TypeId {
		topicType := topicModel.TopicType{
			TopicId:  createData.ID,
			TypeId:   typeItem,
			CreateAt: createTime,
		}
		createTopicType = append(createTopicType, topicType)
	}
	TTypeErr := midTopicTypeCreate.CreateMultipleService(createTopicType, tx)
	if TTypeErr != nil {
		tx.Rollback()
		res.Code = e.CREATE_DATA_FALE
	}

	commitErr := tx.Commit().Error
	if commitErr != nil {
		res.Code = e.CREATE_DATA_FALE
	}

	return res
}
