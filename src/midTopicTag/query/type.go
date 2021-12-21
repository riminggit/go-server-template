package midTopicTagQuery

import (
	"go-server-template/model/topic"
)

type queryTopicTagMidParams struct {
	Id       string   `json:"id"`
	TopicId  string   `json:"topic_id"`
	TagId    string   `json:"tag_id"`
	CreateAt []string `json:"create_at"`
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
	IsUse    string   `json:"is_use"`
}

type queryTopicTagMidReturn struct {
	Code int                   `json:"code"`
	Data []topicModel.TopicTag `json:"data"`
}
