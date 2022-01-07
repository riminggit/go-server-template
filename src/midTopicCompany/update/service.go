package midTopicCompanyDelete

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicCompany/create"
	"go-server-template/src/midTopicCompany/helper"
	"time"
	"gorm.io/gorm"
)

func UpdateService(params UpdateParams, db *gorm.DB) error {

	var queryInfo []topicModel.TopicCompany

	queryFun := db
	queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	queryFun = queryFun.Where("company_id = ?", params.CompanyId)
	queryFun.Model(&topicModel.TopicCompany{}).Find(&queryInfo)

	if len(queryInfo) > 0 {
		setData := topicModel.TopicCompany{
			UpdateAt:  time.Now().Add(8 * time.Hour),
			CompanyId: params.NewCompanyId,
		}
		res := queryFun.Updates(setData)
		thisHelper.CleanRedisQuery()
		return res.Error
	} else {
		createData := topicModel.TopicCompany{
			CompanyId: params.NewCompanyId,
			TopicId:   params.TopicId,
		}
		return midTopicCompanyCreate.CreateService(createData, db)
		
	}
}

