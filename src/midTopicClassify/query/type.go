package midTopicClassifyQuery

import (
	"go-server-template/model/topic"
)

type QueryTopicClassifyMidParams struct {
	Id         string   `json:"id"`
	TopicId    string   `json:"topic_id"`
	ClassifyId string   `json:"classify_id"`
	CreateAt   []string `json:"create_at"`
	DeleteAt   []string `json:"delete_at"`
	UpdateAt   []string `json:"update_at"`
}

type QueryTopicClassifyMidReturn struct {
	Code int                        `json:"code"`
	Data []topicModel.TopicClassify `json:"data"`
}

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type QueryTopicClassifyMidPadingParams struct {
	TopicId    []string `json:"topic_id"`
	ClassifyId []string `json:"classify_id"`
	CreateAt   []string `json:"create_at"`
	DeleteAt   []string `json:"delete_at"`
	UpdateAt   []string `json:"update_at"`
	PageNum    int      `json:"pageNum"`
	PageSize   int      `json:"pageSize"`
}

type QueryTopicClassifyMidPadingReturn struct {
	Code int                                   `json:"code"`
	Data QueryTopicClassifyMidPadingReturnData `json:"data"`
}

type QueryTopicClassifyMidPadingReturnData struct {
	Data           []topicModel.TopicClassify `json:"data"`
	PagingArgument PageArgument               `json:"pageArgument"`
}
