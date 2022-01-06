package midTopicTypeCreate

type CreateParams struct {
	TopicId    int `json:"topic_id"`
	TypeId int `json:"type_id"`
}

type CreateParamsMultiple struct {
	Data []CreateParams
}

type CreateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}