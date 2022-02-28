package midTopicClassifyUpdate

type UpdateParams struct {
	TopicId    int   `json:"topic_id"`
	ClassifyId []int `json:"classify_id"`
	UpdateAt   int64
}
