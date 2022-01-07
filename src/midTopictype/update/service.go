package midTopicTypeUpdate

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicType/create"
	"go-server-template/src/midTopicType/delete"
	"gorm.io/gorm"
	"strconv"
)

func UpdateService(params UpdateParams, tx *gorm.DB) error {
	var Err error

	typeParams := midTopicTypeDelete.DeleteParams{
		TopicId: strconv.Itoa(params.TopicId),
	}
	delTypeErr := midTopicTypeDelete.DeleteService(typeParams, tx)
	if delTypeErr != nil {
		return delTypeErr
	}

	var createTopicType []topicModel.TopicType
	for _, typeItem := range params.TypeId {
		topicType := topicModel.TopicType{
			TopicId:  params.TopicId,
			TypeId:   typeItem,
			CreateAt: params.UpdateAt,
		}
		createTopicType = append(createTopicType, topicType)
	}
	TCErr := midTopicTypeCreate.CreateMultipleService(createTopicType, tx)
	if TCErr != nil {
		return delTypeErr
	}

	return Err
}
