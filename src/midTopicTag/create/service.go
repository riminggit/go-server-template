package midTopicTagCreate

import (
	"go-server-template/model/topic"
	"gorm.io/gorm"
	"time"
)

func CreateService(params CreateParams, db *gorm.DB) error {
	createData := topicModel.TopicTag{
		TopicId:   params.TopicId,
		TagId: params.TagId,
		CreateAt:  time.Now().Add(8 * time.Hour),
	}
	return db.Model(&topicModel.TopicTag{}).Create(createData).Error
}

func CreateMultipleService(params CreateParamsMultiple, db *gorm.DB) error {

	createData := []topicModel.TopicTag{}
	for _, item := range params.Data {
		setData := topicModel.TopicTag{
			TopicId:   item.TopicId,
			TagId: item.TagId,
			CreateAt:  time.Now().Add(8 * time.Hour),
		}
		createData = append(createData, setData)
	}
	return db.Model(&topicModel.TopicTag{}).Create(createData).Error
	
}
