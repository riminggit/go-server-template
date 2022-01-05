package midTopicCompanyDelete


type UpdateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
type UpdateParams struct {
	TopicId       int    `json:"topic_id"`
	CompanyId    string `json:"company_id"`
	NewCompanyId int    `json:"new_company_id"`
}

type UpdateMultiple struct {
	Data []UpdateParams
}
