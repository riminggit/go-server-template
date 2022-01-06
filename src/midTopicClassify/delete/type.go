package midTopicClassifyDelete

type DeleteParams struct {
	ID         string `json:"id"`
	TopicId    string `json:"topic_id"`
	ClassifyId string `json:"classify_id"`
}

type DeleteMultiple struct {
	IDList         []string `json:"id_list"`
	TopicIdList    []string `json:"topic_id_list"`
}
