package midTopicTagUpdate

type UpdateParams struct {
	TopicId  int    `json:"topic_id"`
	TagId    string `json:"tag_id"`
	NewTagId int    `json:"new_tag_id"`
}

type UpdateMultiple struct {
	Data []UpdateParams
}
