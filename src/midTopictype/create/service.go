package midTopicTypeCreate

import (
	"go-server-template/model/topic"
	"gorm.io/gorm"
	"time"
)

func CreateService(params CreateParams, db *gorm.DB) error {

	createData := topicModel.TopicType{
		TopicId:   params.TopicId,
		TypeId: params.TypeId,
		CreateAt:  time.Now().Add(8 * time.Hour),
	}

	return db.Model(&topicModel.TopicType{}).Create(createData).Error
	
}

func CreateMultipleService(params CreateParamsMultiple, db *gorm.DB) error {
	createData := []topicModel.TopicType{}
	for _, item := range params.Data {
		setData := topicModel.TopicType{
			TopicId:   item.TopicId,
			TypeId: item.TypeId,
			CreateAt:  time.Now().Add(8 * time.Hour),
		}
		createData = append(createData, setData)
	}
	return db.Model(&topicModel.TopicType{}).Create(createData).Error
}
