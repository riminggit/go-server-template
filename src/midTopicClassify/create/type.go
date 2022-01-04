package midTopicClassifyCreate

type CreateParams struct {
	TopicId    int `json:"topic_id"`
	ClassifyId int `json:"classify_id"`
}

type CreateParamsMultiple struct {
	Data []CreateParams
}

type CreateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}