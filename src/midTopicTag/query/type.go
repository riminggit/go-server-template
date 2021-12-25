package midTopicTagQuery

import (
	"go-server-template/model/topic"
)

type QueryTopicTagMidParams struct {
	Id       string   `json:"id"`
	TopicId  string   `json:"topic_id"`
	TagId    string   `json:"tag_id"`
	CreateAt []string `json:"create_at"`
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
	IsUse    string   `json:"is_use"`
}

type QueryTopicTagMidReturn struct {
	Code int                   `json:"code"`
	Data []topicModel.TopicTag `json:"data"`
}

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type QueryTopicTagMidPadingParams struct {
	Id       string   `json:"id"`
	TopicId  []string   `json:"topic_id"`
	TagId    []string   `json:"tag_id"`
	CreateAt []string `json:"create_at"`
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
	IsUse    string   `json:"is_use"`
	PageNum  int      `json:"pageNum"`
	PageSize int      `json:"pageSize"`
}

type QueryTopicTagMidPadingReturn struct {
	Code int                              `json:"code"`
	Data QueryTopicTagMidPadingReturnData `json:"data"`
}

type QueryTopicTagMidPadingReturnData struct {
	Data           []topicModel.TopicTag `json:"data"`
	PagingArgument PageArgument          `json:"pageArgument"`
}
