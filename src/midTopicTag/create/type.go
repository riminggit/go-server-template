package midTopicTagCreate

type CreateParams struct {
	TopicId int `json:"topic_id"`
	TagId   int `json:"tag_id"`
}

type CreateParamsMultiple struct {
	Data []CreateParams
}

type CreateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
