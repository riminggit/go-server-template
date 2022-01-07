package midTopicTypeCreate

import (
	"go-server-template/model/topic"
	"gorm.io/gorm"
)

func CreateService(params topicModel.TopicType, db *gorm.DB) error {
	return db.Model(&topicModel.TopicType{}).Create(params).Error
}

func CreateMultipleService(params []topicModel.TopicType, db *gorm.DB) error {
	return db.Model(&topicModel.TopicType{}).Create(params).Error
}
