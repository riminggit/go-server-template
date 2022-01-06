package midTopicTagDelete

type DeleteParams struct {
	ID      string `json:"id"`
	TopicId string `json:"topic_id"`
	TagId   string `json:"tag_id"`
}

type DeleteMultiple struct {
	IDList      []string `json:"id_list"`
	TopicIdList []string `json:"topic_id_list"`
	TagIdList   []string `json:"tag_id_list"`
}

type DeleteReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
