package midTopicTypeDelete

type DeleteParams struct {
	ID      string `json:"id"`
	TopicId string `json:"topic_id"`
	TypeId  string `json:"type_id"`
}

type DeleteMultiple struct {
	IDList      []string `json:"id_list"`
	TopicIdList []string `json:"topic_id_list"`
	TypeIdList  []string `json:"type_id_list"`
}

type DeleteReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
