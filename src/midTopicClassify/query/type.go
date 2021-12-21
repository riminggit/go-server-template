package midTopicClassifyQuery

import (
	"go-server-template/model/topic"
)

type queryTopicClassifyMidParams struct {
	Id         string   `json:"id"`
	TopicId    string   `json:"topic_id"`
	ClassifyId string   `json:"classify_id"`
	CreateAt   []string `json:"create_at"`
	DeleteAt   []string `json:"delete_at"`
	UpdateAt   []string `json:"update_at"`
	IsUse      string   `json:"is_use"`
}

type queryTopicClassifyMidReturn struct {
	Code int                        `json:"code"`
	Data []topicModel.TopicClassify `json:"data"`
}
