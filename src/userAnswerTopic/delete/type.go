package userAnswerTopicDelete

type DeleteParams struct {
	TopicIdList      string `json:"topic_id_list"`       // 题目id列表
	TopicSetId       int    `json:"topic_set_id"`        // 套题id
}

type DeleteReturn struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
