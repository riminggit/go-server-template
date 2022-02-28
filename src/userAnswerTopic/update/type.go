package userAnswerTopicUpdate

// 判断由前端处理
type UserAnswerTopicUpdateParams struct {
	UserId           int    `json:"user_id"`             // 用户id
	TopicIdList      string `json:"topic_id_list"`       // 题目id列表
	TopicSetId       int    `json:"topic_set_id"`        // 套题id
	AnswerEnd        int64  `json:"answer_end"`          // 答题结束时间
	AnswerCorrectNum int    `json:"answer_correct_num" ` // 答题正确数量
	IsAchieve        int    `json:"is_achieve" `         // 是否完成答题 0否  1 完成
	Remark           string `json:"remark" `             // 备注
	UpdateAt         int64  `json:"update_at"`           // 创建时间
	IsUse            int    `json:"is_use" `             // 是否删除:0 删除 1未删除
}

type UpdateReturn struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
