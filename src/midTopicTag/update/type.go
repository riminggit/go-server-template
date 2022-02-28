package midTopicTagUpdate

type UpdateParams struct {
	TopicId  int   `json:"topic_id"`
	TagId    []int `json:"tag_id"`
	UpdateAt int64
}
