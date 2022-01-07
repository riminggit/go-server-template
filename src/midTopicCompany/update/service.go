package midTopicCompanyUpdate

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicCompany/create"
	"go-server-template/src/midTopicCompany/delete"
	"gorm.io/gorm"
	"strconv"
)

func UpdateService(params UpdateParams, tx *gorm.DB) error {
	var Err error

	companyParams := midTopicCompanyDelete.DeleteParams{
		TopicId: strconv.Itoa(params.TopicId),
	}
	delCompanyErr := midTopicCompanyDelete.DeleteService(companyParams, tx)
	if delCompanyErr != nil {
		return delCompanyErr
	}

	var createTopicCompany []topicModel.TopicCompany
	for _, companyItem := range params.CompanyId {
		topicCompany := topicModel.TopicCompany{
			TopicId:   params.TopicId,
			CompanyId: companyItem,
			CreateAt:  params.UpdateAt,
		}
		createTopicCompany = append(createTopicCompany, topicCompany)
	}
	TCErr := midTopicCompanyCreate.CreateMultipleService(createTopicCompany, tx)
	if TCErr != nil {
		return delCompanyErr
	}

	return Err
}
