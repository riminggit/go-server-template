package midTopicCompanyDelete

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicCompany/helper"
	"gorm.io/gorm"
)

func DeleteService(params DeleteParams, db *gorm.DB) error {
	var delErr error

	if params.ID != "" {
		delErr = db.Delete(&topicModel.TopicCompany{}, "id = ?", params.ID).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}

	if params.TopicId != "" && params.CompanyId == "" {
		delErr = db.Delete(&topicModel.TopicCompany{}, "topic_id = ?", params.TopicId).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}

	if params.TopicId != "" && params.CompanyId != "" {
		delErr = db.Where("topic_id = ?", params.TopicId).Delete(&topicModel.TopicCompany{}, "company_id = ?", params.CompanyId).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}
	return delErr
}

func DeleteMultipleService(params DeleteMultiple, db *gorm.DB) error {
	var delErr error
	if len(params.IDList) > 0 {
		delErr = db.Delete(&topicModel.TopicCompany{}, "id IN ?", params.IDList).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}
	if len(params.TopicIdList) > 0 {
		delErr = db.Delete(&topicModel.TopicCompany{}, "topic_id IN ?", params.TopicIdList).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}

	return delErr
}
