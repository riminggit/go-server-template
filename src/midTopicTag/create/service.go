package midTopicTagCreate

import (
	"go-server-template/model/topic"
	"gorm.io/gorm"
)

func CreateService(params topicModel.TopicTag, db *gorm.DB) error {
	return db.Model(&topicModel.TopicTag{}).Create(params).Error
}

func CreateMultipleService(params []topicModel.TopicTag, db *gorm.DB) error {
	return db.Model(&topicModel.TopicTag{}).Create(params).Error
}
