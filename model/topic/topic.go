package topicModel

type Topic struct {
	ID               int    `json:"id" gorm:"type:bigint;column:id;comment:id;not null"` // id
	Title            string `json:"title" gorm:"type:varchar(255);not null;unique"`      // 题目标题
	QuestionType     int    `json:"question_type" gorm:"type:int(3);"`                   // 题目类型，1为选择题，2为解析题，3既是选择又是解析
	Analysis         string `json:"analysis" gorm:"type:text"`                           // 题解，内容为富文本
	SelectAnalysis   string `json:"select_analysis" gorm:"type:text"`                    // 选择题答案，对象字符串
	Degree           int    `json:"degree" gorm:"type:int(3);"`                          // 难度，简单：0，中等：1，难：2，极难：3
	Level            int    `json:"level" gorm:"type:int(3);"`                           // 1 初级 ， 2 中级 ， 3 高级 ， 4 资深 ，5 专家 ， 6 资深专家 ， 7 研究员
	IsBaseTopic      int    `json:"is_base_topic" gorm:"type:int(3);"`                   // 0 否  1是 ，是否是基础题
	IsImportantTopic int    `json:"is_important_topic" gorm:"type:int(3);"`              // 0 否  1是 ，是否是重点题
	ComeFrom         string `json:"come_from" gorm:"type:varchar(255);not null"`         // 来源
	CreateAt         int64  `json:"create_at"`                                           // 创建时间
	DeleteAt         int64  `json:"delete_at"`
	UpdateAt         int64  `json:"update_at"`
	IsUse            int    `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}
