package midTopicClassifyUpdate

type UpdateParams struct {
	TopicId       int    `json:"topic_id"`
	ClassifyId    string `json:"classify_id"`
	NewClassifyId int    `json:"new_classify_id"`
}

