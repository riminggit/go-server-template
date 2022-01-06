package midTopicTypeDelete

import (
	"go-server-template/model/topic"
	"gorm.io/gorm"
	"go-server-template/src/midTopicType/helper"
)

func DeleteService(params DeleteParams, db *gorm.DB) error {
	var delErr error
	if params.ID != "" {
		delErr = db.Delete(&topicModel.TopicType{}, "id = ?", params.ID).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}

	if params.TopicId != "" && params.TypeId == ""{
		delErr = db.Delete(&topicModel.TopicType{}, "topic_id = ?", params.TopicId).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}

	if  params.TopicId != "" && params.TypeId != "" {
		delErr = db.Where("topic_id = ?", params.TopicId).Delete(&topicModel.TopicType{}, "type_id = ?", params.TypeId).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}
	return delErr
}

func DeleteMultipleService(params DeleteMultiple, db *gorm.DB) error {
	var delErr error

	if len(params.IDList) > 0 {
		delErr = db.Delete(&topicModel.TopicType{}, "id IN ?", params.IDList).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}
	if len(params.TopicIdList) > 0 {
		delErr = db.Delete(&topicModel.TopicType{}, "topic_id IN ?", params.TopicIdList).Error
		thisHelper.CleanRedisQuery()
		return delErr
	}
	return delErr
}
