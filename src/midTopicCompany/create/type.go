package midTopicCompanyCreate

type CreateParams struct {
	TopicId    int `json:"topic_id"`
	CompanyId int `json:"company_id"`
}

type CreateParamsMultiple struct {
	Data []CreateParams
}

type CreateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}