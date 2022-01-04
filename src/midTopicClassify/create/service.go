package midTopicClassifyCreate

import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"time"
)

func CreateService(params CreateParams) *CreateReturn {
	res := &CreateReturn{}
	createData := topicModel.TopicClassify{
		TopicId:    params.TopicId,
		ClassifyId: params.ClassifyId,
		CreateAt:   time.Now().Add(8 * time.Hour),
	}

	err := DB.DBLivingExample.Model(&topicModel.TopicClassify{}).Create(createData).Error
	if err != nil {
		res.Code = e.CREATE_DATA_FILE
		res.Data = append(res.Data, "新建题目-分类关联失败")
		return res
	}
	res.Code = e.SUCCESS
	return res
}


func CreateMultipleService(params CreateParamsMultiple) *CreateReturn {
	res := &CreateReturn{}
	createData := []topicModel.TopicClassify{}
	for _, item := range params.Data {
		setData := topicModel.TopicClassify{
			TopicId:    item.TopicId,
			ClassifyId: item.ClassifyId,
			CreateAt:   time.Now().Add(8 * time.Hour),
		}
		createData = append(createData, setData)
	}
	err := DB.DBLivingExample.Model(&topicModel.TopicClassify{}).Create(createData).Error
	if err != nil {
		res.Code = e.CREATE_DATA_FILE
		res.Data = append(res.Data, "新建题目-分类关联失败")
		return res
	}
	res.Code = e.SUCCESS
	return res
}
