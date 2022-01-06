package midTopicClassifyCreate

type CreateParams struct {
	TopicId    int `json:"topic_id"`
	ClassifyId int `json:"classify_id"`
}

type CreateParamsMultiple struct {
	Data []CreateParams
}