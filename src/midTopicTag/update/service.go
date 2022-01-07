package midTopicTagUpdate

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicTag/create"
	"go-server-template/src/midTopicTag/delete"
	"gorm.io/gorm"
	"strconv"
)

func UpdateService(params UpdateParams, tx *gorm.DB) error {
	var Err error

	tagParams := midTopicTagDelete.DeleteParams{
		TopicId: strconv.Itoa(params.TopicId),
	}
	delTagErr := midTopicTagDelete.DeleteService(tagParams, tx)
	if delTagErr != nil {
		return delTagErr
	}

	var createTopicTag []topicModel.TopicTag
	for _, tagItem := range params.TagId {
		topicTag := topicModel.TopicTag{
			TopicId:    params.TopicId,
			TagId: tagItem,
			CreateAt:   params.UpdateAt,
		}
		createTopicTag = append(createTopicTag, topicTag)
	}
	TCErr := midTopicTagCreate.CreateMultipleService(createTopicTag, tx)
	if TCErr != nil {
		return delTagErr
	}

	return Err
}
