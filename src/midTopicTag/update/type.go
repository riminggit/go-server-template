package midTopicTagUpdate

import "time"

type UpdateParams struct {
	TopicId  int   `json:"topic_id"`
	TagId    []int `json:"tag_id"`
	UpdateAt time.Time
}
