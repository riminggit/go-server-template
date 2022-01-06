package midTopicClassifyCreate

import (
	"go-server-template/model/topic"
	"gorm.io/gorm"
	"time"
)

func CreateService(params CreateParams, db *gorm.DB) error {

	createData := topicModel.TopicClassify{
		TopicId:    params.TopicId,
		ClassifyId: params.ClassifyId,
		CreateAt:   time.Now().Add(8 * time.Hour),
	}

	return db.Model(&topicModel.TopicClassify{}).Create(createData).Error

}

func CreateMultipleService(params CreateParamsMultiple, db *gorm.DB) error {
	createData := []topicModel.TopicClassify{}
	for _, item := range params.Data {
		setData := topicModel.TopicClassify{
			TopicId:    item.TopicId,
			ClassifyId: item.ClassifyId,
			CreateAt:   time.Now().Add(8 * time.Hour),
		}
		createData = append(createData, setData)
	}
	return db.Model(&topicModel.TopicClassify{}).Create(createData).Error

}
