package topicSetQuery

import (
	"go-server-template/model/topic"
)

type QueryTopicSetParams struct {
	Id                 []string `json:"id"`
	Name               string   `json:"name"`
	TopicSetIdList     string   `json:"topic_set_id_list"`
	TopicSetDifficulty string   `json:"topic_set_difficulty"`
	TopicSetLevel      string   `json:"topic_set_level"`
	TopicType          string   `json:"topic_type"` // 练习：1  考试 ：2
	CreateAt           []string `json:"create_at"`
	DeleteAt           []string `json:"delete_at"`
	UpdateAt           []string `json:"update_at"`
	PageNum            int      `json:"pageNum"`
	PageSize           int      `json:"pageSize"`
	NeedTopicDetail    string   `json:"need_topic_detail"`
	IsUse              string   `json:"is_use"`
}
type queryTopicSetReturn struct {
	Code int                     `json:"code"`
	Data queryTopicSetReturnData `json:"data"`
}

type queryTopicSetReturnData struct {
	Data           []topicModel.TopicSet `json:"data"`
	PagingArgument PageArgument          `json:"pageArgument"`
}

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type queryTopicSetSimpleReturn struct {
	Code int                   `json:"code"`
	Data []topicModel.TopicSet `json:"data"`
}
