package midTopicClassifyUpdate

type UpdateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
type UpdateParams struct {
	TopicId       int    `json:"topic_id"`
	ClassifyId    string `json:"classify_id"`
	NewClassifyId int    `json:"new_classify_id"`
}

type UpdateMultiple struct {
	Data []UpdateParams
}
