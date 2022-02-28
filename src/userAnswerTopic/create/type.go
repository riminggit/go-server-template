package userAnswerTopicCreate

// 判断由前端处理
type UserAnswerTopicCreateParams struct {
	UserId          int    `json:"user_id" `          // 用户id
	TopicIdList     string `json:"topic_id_list"`     // 题目id列表
	TopicSetId      int    `json:"topic_set_id"`      // 套题id
	AnswerStart     string `json:"answer_start"`      // 答题开始时间
	AnswerNum       int    `json:"answer_num" `       // 答题数量
	IsAchieve       int    `json:"is_achieve"`        // 是否完成答题 0否  1 完成
	TopicDifficulty int    `json:"topic_difficulty" ` // 题目难度，由所有的题目难度指数相加后除题目数
	TopicLevel      int    `json:"topic_level" `      // 题目等级，由所有的题目等级指数相加后除题目数
	Remark          string `json:"remark"`            // 备注
	CreateAt        int64  `json:"create_at"`         // 创建时间
	IsUse           int    `json:"is_use" `           // 是否删除:0 删除 1未删除
}

type CreateReturn struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
