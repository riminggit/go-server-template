package topicModel

import (
	"time"
)

type Topic struct {
	ID               int       `gorm:"autoIncrement;primaryKey;type:int(11);uniqueIndex;not null"` // id
	Title            string    `gorm:"type:varchar(255);not null"`                                 // 题目标题
	QuestionType     int       `gorm:"type:int(3);"`                                               // 题目类型，1为选择题，2为解析题，3既是选择又是解析
	Analysis         string    `gorm:"type:text"`                                                  // 题解，内容为富文本
	SelectAnalysis   string    `gorm:"type:text"`                                                  // 选择题答案，对象字符串
	KeywordList      string    `gorm:"type:varchar(255);"`                                         // 关键字数组，存题目相关的关键字
	Degree           int       `gorm:"type:int(3);"`                                               // 难度，简单：0，中等：1，难：2，极难：3
	Level            int       `gorm:"type:int(3);"`                                               // 1 初级 ， 2 中级 ， 3 高级 ， 4 资深 ，5 专家 ， 6 资深专家 ， 7 研究员
	IsBaseTopic      int       `gorm:"type:int(3);"`                                               // 0 否  1是 ，是否是基础题
	IsImportantTopic int       `gorm:"type:int(3);"`                                               // 0 否  1是 ，是否是重点题
	ComeFrom         string    `gorm:"type:varchar(255);not null"`                                 // 来源
	CreateAt         time.Time // 创建时间
	IsUse            int       `gorm:"type:int(3);not null"` // 是否删除:0 删除 1未删除
}
