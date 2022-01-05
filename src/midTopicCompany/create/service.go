package midTopicCompanyCreate

import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"time"
)

func CreateService(params CreateParams) *CreateReturn {
	res := &CreateReturn{}
	createData := topicModel.TopicCompany{
		TopicId:   params.TopicId,
		CompanyId: params.CompanyId,
		CreateAt:  time.Now().Add(8 * time.Hour),
	}

	err := DB.DBLivingExample.Model(&topicModel.TopicCompany{}).Create(createData).Error
	if err != nil {
		res.Code = e.CREATE_DATA_FILE
		res.Data = append(res.Data, "新建题目-公司关联失败")
		return res
	}
	res.Code = e.SUCCESS
	return res
}

func CreateMultipleService(params CreateParamsMultiple) *CreateReturn {
	res := &CreateReturn{}
	createData := []topicModel.TopicCompany{}
	for _, item := range params.Data {
		setData := topicModel.TopicCompany{
			TopicId:   item.TopicId,
			CompanyId: item.CompanyId,
			CreateAt:  time.Now().Add(8 * time.Hour),
		}
		createData = append(createData, setData)
	}
	err := DB.DBLivingExample.Model(&topicModel.TopicCompany{}).Create(createData).Error
	if err != nil {
		res.Code = e.CREATE_DATA_FILE
		res.Data = append(res.Data, "新建题目-分类关联失败")
		return res
	}
	res.Code = e.SUCCESS
	return res
}
