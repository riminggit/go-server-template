package midTopicCompanyUpdate

type UpdateParams struct {
	TopicId   int   `json:"topic_id"`
	CompanyId []int `json:"company_id"`
	UpdateAt  int64
}
