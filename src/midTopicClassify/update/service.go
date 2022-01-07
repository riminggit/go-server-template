package midTopicClassifyUpdate

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicClassify/create"
	"go-server-template/src/midTopicClassify/delete"
	"gorm.io/gorm"
	"strconv"
)

func UpdateService(params UpdateParams, tx *gorm.DB) error {
	var Err error

	classifyParams := midTopicClassifyDelete.DeleteParams{
		TopicId: strconv.Itoa(params.TopicId),
	}
	delClassifyErr := midTopicClassifyDelete.DeleteService(classifyParams, tx)
	if delClassifyErr != nil {
		return delClassifyErr
	}

	var createTopicClassify []topicModel.TopicClassify
	for _, classifyItem := range params.ClassifyId {
		topicClassify := topicModel.TopicClassify{
			TopicId:    params.TopicId,
			ClassifyId: classifyItem,
			CreateAt:   params.UpdateAt,
		}
		createTopicClassify = append(createTopicClassify, topicClassify)
	}
	TCErr := midTopicClassifyCreate.CreateMultipleService(createTopicClassify, tx)
	if TCErr != nil {
		return delClassifyErr
	}

	return Err
}
