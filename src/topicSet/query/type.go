package topicSetQuery

import (
	"go-server-template/model/topic"
)


type QueryTopicSetParams struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	TopicSetIdList     string   `json:"topic_set_id_list"`
	TopicSetDifficulty string   `json:"topic_set_difficulty"`
	TopicSetLevel      string   `json:"topic_set_level"`
	CreateAt           []string `json:"create_at"`
	DeleteAt           []string `json:"delete_at"`
	UpdateAt           []string `json:"update_at"`
	IsUse              string   `json:"is_use"`
}
type queryTopicSetReturn struct {
	Code int                `json:"code"`
	Data []topicModel.TopicSet `json:"data"`
}
