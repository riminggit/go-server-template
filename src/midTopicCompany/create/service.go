package midTopicCompanyCreate

import (
	"go-server-template/model/topic"
	"gorm.io/gorm"
	"time"
)

func CreateService(params CreateParams, db *gorm.DB) error {
	createData := topicModel.TopicCompany{
		TopicId:   params.TopicId,
		CompanyId: params.CompanyId,
		CreateAt:  time.Now().Add(8 * time.Hour),
	}

	return db.Model(&topicModel.TopicCompany{}).Create(createData).Error
	
}

func CreateMultipleService(params CreateParamsMultiple, db *gorm.DB) error {

	createData := []topicModel.TopicCompany{}
	for _, item := range params.Data {
		setData := topicModel.TopicCompany{
			TopicId:   item.TopicId,
			CompanyId: item.CompanyId,
			CreateAt:  time.Now().Add(8 * time.Hour),
		}
		createData = append(createData, setData)
	}
	return db.Model(&topicModel.TopicCompany{}).Create(createData).Error
}
