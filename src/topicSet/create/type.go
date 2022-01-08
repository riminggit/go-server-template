package topicSetCreate

type CreateParams struct {
	Name               string   `json:"name"`                 // 套题名称
	TopicSetIdList     []string `json:"topic_set_id_list"`    // 题目id列表
	TopicSetDifficulty int      `json:"topic_set_difficulty"` // 题目难度，由所有的题目难度指数相加后除题目数
	TopicSetLevel      int      `json:"topic_set_level" `     // 题目等级，由所有的题目等级指数相加后除题目数
	Remark             string   `json:"remark"`               // 备注
}

type CreateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
