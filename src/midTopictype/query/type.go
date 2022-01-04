package midTopicTypeQuery

import (
	"go-server-template/model/topic"
)

type QueryTopicTypeMidParams struct {
	Id       string   `json:"id"`
	TopicId  string   `json:"topic_id"`
	TypeId   string   `json:"type_id"`
	CreateAt []string `json:"create_at"`
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
}

type QueryTopicTypeMidReturn struct {
	Code int                    `json:"code"`
	Data []topicModel.TopicType `json:"data"`
}

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type QueryTopicTypeMidPadingParams struct {
	Id       string   `json:"id"`
	TopicId  []string `json:"topic_id"`
	TypeId   []string   `json:"type_id"`
	CreateAt []string `json:"create_at"`
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
	PageNum  int      `json:"pageNum"`
	PageSize int      `json:"pageSize"`
}

type QueryTopicTypeMidPadingReturn struct {
	Code int                               `json:"code"`
	Data QueryTopicTypeMidPadingReturnData `json:"data"`
}

type QueryTopicTypeMidPadingReturnData struct {
	Data           []topicModel.TopicType `json:"data"`
	PagingArgument PageArgument           `json:"pageArgument"`
}
