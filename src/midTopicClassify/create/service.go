package midTopicClassifyCreate

import (
	"go-server-template/model/topic"
	"gorm.io/gorm"
)

func CreateService(params topicModel.TopicClassify, db *gorm.DB) error {
	return db.Model(&topicModel.TopicClassify{}).Create(params).Error
}

func CreateMultipleService(params []topicModel.TopicClassify, db *gorm.DB) error {
	return db.Model(&topicModel.TopicClassify{}).Create(&params).Error
}
