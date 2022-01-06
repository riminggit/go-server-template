package midTopicTypeUpdate

type UpdateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
type UpdateParams struct {
	TopicId   int    `json:"topic_id"`
	TypeId    string `json:"type_id"`
	NewTypeId int    `json:"new_type_id"`
}

type UpdateMultiple struct {
	Data []UpdateParams
}
