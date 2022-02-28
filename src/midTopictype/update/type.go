package midTopicTypeUpdate

type UpdateParams struct {
	TopicId  int   `json:"topic_id"`
	TypeId   []int `json:"type_id"`
	UpdateAt int64
}
