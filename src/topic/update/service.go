package topicUpdate

import (
	"github.com/gin-gonic/gin"
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/midTopicClassify/update"
	"go-server-template/src/midTopicCompany/update"
	"go-server-template/src/midTopicTag/update"
	"go-server-template/src/midTopicType/update"
	"time"
)

func UpdateService(c *gin.Context, params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}
	res.Code = e.SUCCESS

	tx := DB.DBLivingExample.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			res.Code = e.CREATE_DATA_FILE
			res.Data = append(res.Data, "题目修改失败")
		}
	}()

	updateTime := time.Now().Add(8 * time.Hour)

	updateData := &topicModel.Topic{
		Degree:           params.Degree,
		IsBaseTopic:      params.IsBaseTopic,
		IsImportantTopic: params.IsImportantTopic,
		UpdateAt:         updateTime,
	}

	if params.Title != "" {
		updateData.Title = params.Title
	}
	if params.QuestionType != 0 {
		updateData.QuestionType = params.QuestionType
	}
	if params.Analysis != "" {
		updateData.Analysis = params.Analysis
	}
	if params.SelectAnalysis != "" {
		updateData.SelectAnalysis = params.SelectAnalysis
	}
	if params.Level != 0 {
		updateData.Level = params.Level
	}
	if params.ComeFrom != "" {
		updateData.ComeFrom = params.ComeFrom
	}

	err := tx.Model(&topicModel.Topic{}).Where("id = ?", params.ID).Updates(updateData).Error
	if err != nil {
		tx.Rollback()
		res.Code = e.CREATE_DATA_FILE
	}

	classifyUpdate := midTopicClassifyUpdate.UpdateParams{
		TopicId:    params.ID,
		ClassifyId: params.ClassifyId,
		UpdateAt:   updateTime,
	}
	TCErr := midTopicClassifyUpdate.UpdateService(classifyUpdate, tx)
	if TCErr != nil {
		tx.Rollback()
		res.Code = e.UPDATE_DATA_FILE
	}

	companyUpdate := midTopicCompanyUpdate.UpdateParams{
		TopicId:   params.ID,
		CompanyId: params.CompanyId,
		UpdateAt:  updateTime,
	}
	TCompanyErr := midTopicCompanyUpdate.UpdateService(companyUpdate, tx)
	if TCompanyErr != nil {
		tx.Rollback()
		res.Code = e.UPDATE_DATA_FILE
	}

	tagUpdate := midTopicTagUpdate.UpdateParams{
		TopicId:  params.ID,
		TagId:    params.TagId,
		UpdateAt: updateTime,
	}
	TTagErr := midTopicTagUpdate.UpdateService(tagUpdate, tx)
	if TTagErr != nil {
		tx.Rollback()
		res.Code = e.UPDATE_DATA_FILE
	}

	typeUpdate := midTopicTypeUpdate.UpdateParams{
		TopicId:  params.ID,
		TypeId:   params.TypeId,
		UpdateAt: updateTime,
	}
	TTypeErr := midTopicTypeUpdate.UpdateService(typeUpdate, tx)
	if TTypeErr != nil {
		tx.Rollback()
		res.Code = e.UPDATE_DATA_FILE
	}

	commitErr := tx.Commit().Error
	if commitErr != nil {
		res.Code = e.CREATE_DATA_FILE
	}

	return res
}
