package midTopicCompanyQuery

import (
	"go-server-template/model/topic"
)

type QueryTopicCompanyMidParams struct {
	Id        string   `json:"id"`
	TopicId   string   `json:"topic_id"`
	TopicTime string   `json:"topic_time"`
	CompanyId string   `json:"company_id"`
	CreateAt  []string `json:"create_at"`
	DeleteAt  []string `json:"delete_at"`
	UpdateAt  []string `json:"update_at"`
	IsUse     string   `json:"is_use"`
}

type QueryTopicCompanyMidReturn struct {
	Code int                        `json:"code"`
	Data []topicModel.TopicCompany `json:"data"`
}

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type QueryTopicCompanyMidPadingParams struct {
	Id         string   `json:"id"`
	TopicId   []string   `json:"topic_id"`
	TopicTime string   `json:"topic_time"`
	CompanyId []string   `json:"company_id"`
	CreateAt   []string `json:"create_at"`
	DeleteAt   []string `json:"delete_at"`
	UpdateAt   []string `json:"update_at"`
	IsUse      string   `json:"is_use"`
	PageNum    int      `json:"pageNum"`
	PageSize   int      `json:"pageSize"`
}

type QueryTopicCompanyMidPadingReturn struct {
	Code int                                   `json:"code"`
	Data QueryTopicCompanyMidPadingReturnData `json:"data"`
}

type QueryTopicCompanyMidPadingReturnData struct {
	Data           []topicModel.TopicCompany `json:"data"`
	PagingArgument PageArgument               `json:"pageArgument"`
}
