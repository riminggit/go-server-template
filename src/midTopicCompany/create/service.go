package midTopicCompanyCreate

import (
	"go-server-template/model/topic"
	"gorm.io/gorm"
)

func CreateService(params topicModel.TopicCompany, db *gorm.DB) error {
	return db.Model(&topicModel.TopicCompany{}).Create(params).Error
}

func CreateMultipleService(params []topicModel.TopicCompany, db *gorm.DB) error {
	return db.Model(&topicModel.TopicCompany{}).Create(params).Error
}
