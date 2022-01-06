package midTopicClassifyDelete

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicClassify/helper"
	"gorm.io/gorm"
)

func DeleteService(params DeleteParams, db *gorm.DB) error {

	var delErr error
	if params.ID != "" {
		delErr = db.Delete(&topicModel.TopicClassify{}, "id = ?", params.ID).Error
		// 清除查询的redis缓存
		thisHelper.CleanRedisQuery()
		return delErr
	}

	if params.TopicId != "" && params.ClassifyId == "" {
		delErr = db.Delete(&topicModel.TopicClassify{}, "topic_id = ?", params.TopicId).Error
		// 清除查询的redis缓存
		thisHelper.CleanRedisQuery()
		return delErr
	}

	if params.TopicId != "" && params.ClassifyId != "" {
		delErr = db.Where("topic_id = ?", params.TopicId).Delete(&topicModel.TopicClassify{}, "classify_id = ?", params.ClassifyId).Error
		// 清除查询的redis缓存
		thisHelper.CleanRedisQuery()
		return delErr
	}

	return delErr
}

func DeleteMultipleService(params DeleteMultiple, db *gorm.DB) error {
	var delErr error

	if len(params.IDList) > 0 {
		delErr = db.Delete(&topicModel.TopicClassify{}, "id IN ?", params.IDList).Error
		// 清除查询的redis缓存
		thisHelper.CleanRedisQuery()
		return delErr
	}
	if len(params.TopicIdList) > 0  {
		delErr = db.Delete(&topicModel.TopicClassify{}, "topic_id IN ?", params.TopicIdList).Error
		// 清除查询的redis缓存
		thisHelper.CleanRedisQuery()
		return delErr
	}

	return delErr
}
