package midTopicTypeQuery

import (
	"go-server-template/model/topic"
)

type queryTopicTypeMidParams struct {
	Id       string   `json:"id"`
	TopicId  string   `json:"topic_id"`
	TypeId    string   `json:"type_id"`
	CreateAt []string `json:"create_at"`
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
	IsUse    string   `json:"is_use"`
}

type queryTopicTypeMidReturn struct {
	Code int                   `json:"code"`
	Data []topicModel.TopicType `json:"data"`
}
