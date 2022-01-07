package midTopicCompanyUpdate

import "time"
type UpdateParams struct {
	TopicId    int      `json:"topic_id"`
	CompanyId []int `json:"company_id"`
	UpdateAt   time.Time
}
