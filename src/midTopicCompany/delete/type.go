package midTopicCompanyDelete

type DeleteParams struct {
	ID        string `json:"id"`
	TopicId   string `json:"topic_id"`
	CompanyId string    `json:"company_id"`
}

type DeleteMultiple struct {
	IDList        []string `json:"id_list"`
	TopicIdList   []string `json:"topic_id_list"`
	CompanyIdList []string `json:"company_id_list"`
}

type DeleteReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
