package userAnswerTopicUpdate

// 判断由前端处理
type UserAnswerTopicUpdateParams struct {
	UserId           int    `json:"user_id" gorm:"type:bigint;not null"`          // 用户id
	TopicIdList      string `json:"topic_id_list" gorm:"type:text;"`              // 题目id列表
	TopicSetId       string `json:"topic_set_id" gorm:"type:bigint;"`             // 套题id
	AnswerStart      int64  `json:"answer_start"`                                 // 答题开始时间
	AnswerEnd        int64  `json:"answer_end"`                                   // 答题结束时间
	AnswerNum        int    `json:"answer_num" gorm:"type:int(3);"`               // 答题数量
	AnswerCorrectNum int    `json:"answer_correct_num" gorm:"type:int(3);"`       // 答题正确数量
	IsAchieve        int    `json:"is_achieve" gorm:"type:int(3);"`               // 是否完成答题 0否  1 完成
	TopicDifficulty  int    `json:"topic_difficulty" gorm:"type:int(3);"`         // 题目难度，由所有的题目难度指数相加后除题目数
	TopicLevel       int    `json:"topic_level" gorm:"type:int(3);"`              // 题目等级，由所有的题目等级指数相加后除题目数
	Remark           string `json:"remark" gorm:"type:varchar(255);"`             // 备注
	UpdateAt         int64  `json:"update_at"`                                    // 创建时间
	IsUse            int    `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}

type UpdateReturn struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
