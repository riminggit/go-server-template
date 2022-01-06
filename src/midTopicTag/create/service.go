package midTopicTagCreate

import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"time"
)

func CreateService(params CreateParams) *CreateReturn {
	res := &CreateReturn{}
	createData := topicModel.TopicTag{
		TopicId:   params.TopicId,
		TagId: params.TagId,
		CreateAt:  time.Now().Add(8 * time.Hour),
	}

	err := DB.DBLivingExample.Model(&topicModel.TopicTag{}).Create(createData).Error
	if err != nil {
		res.Code = e.CREATE_DATA_FILE
		res.Data = append(res.Data, "新建题目-标签关联失败")
		return res
	}
	res.Code = e.SUCCESS
	return res
}

func CreateMultipleService(params CreateParamsMultiple) *CreateReturn {
	res := &CreateReturn{}
	createData := []topicModel.TopicTag{}
	for _, item := range params.Data {
		setData := topicModel.TopicTag{
			TopicId:   item.TopicId,
			TagId: item.TagId,
			CreateAt:  time.Now().Add(8 * time.Hour),
		}
		createData = append(createData, setData)
	}
	err := DB.DBLivingExample.Model(&topicModel.TopicTag{}).Create(createData).Error
	if err != nil {
		res.Code = e.CREATE_DATA_FILE
		res.Data = append(res.Data, "新建题目-标签关联失败")
		return res
	}
	res.Code = e.SUCCESS
	return res
}
