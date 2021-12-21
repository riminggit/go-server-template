package midTopicCompanyQuery

import (
	"go-server-template/model/topic"
)

type queryTopicCompanyMidParams struct {
	Id        string   `json:"id"`
	TopicId   string   `json:"topic_id"`
	TopicTime string   `json:"topic_time"`
	CompanyId string   `json:"company_id"`
	CreateAt  []string `json:"create_at"`
	DeleteAt  []string `json:"delete_at"`
	UpdateAt  []string `json:"update_at"`
	IsUse     string   `json:"is_use"`
}

type queryTopicCompanyMidReturn struct {
	Code int                        `json:"code"`
	Data []topicModel.TopicCompany `json:"data"`
}
