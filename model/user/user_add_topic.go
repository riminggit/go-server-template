package userModel

import (
	"time"
)

type UserAddTopic struct {
	ID               int       `json:"id" gorm:"column:id;comment:id;not null"`        // id
	UserId           int       `json:"user_id" gorm:"type:bigint;not null"`            // 用户id
	Title            string    `json:"title" gorm:"type:varchar(255);not null;unique"` // 题目标题
	QuestionType     int       `json:"question_type" gorm:"type:int(3);"`              // 题目类型，1为选择题，2为解析题，3既是选择又是解析
	Analysis         string    `json:"analysis" gorm:"type:text"`                      // 题解，内容为富文本
	SelectAnalysis   string    `json:"select_analysis" gorm:"type:text"`               // 选择题答案，对象字符串
	KeywordList      string    `json:"keyword_list" gorm:"type:varchar(255);"`         // 关键字数组，存题目相关的关键字
	Degree           int       `json:"degree" gorm:"type:int(3);"`                     // 难度，简单：0，中等：1，难：2，极难：3
	Level            int       `json:"level" gorm:"type:int(3);"`                      // 1 初级 ， 2 中级 ， 3 高级 ， 4 资深 ，5 专家 ， 6 资深专家 ， 7 研究员
	IsBaseTopic      int       `json:"is_base_topic" gorm:"type:int(3);"`              // 0 否  1是 ，是否是基础题
	IsImportantTopic int       `json:"is_important_topic" gorm:"type:int(3);"`         // 0 否  1是 ，是否是重点题
	ComeFrom         string    `json:"come_from" gorm:"type:varchar(255);not null"`    // 来源
	ClassifyId       int       `json:"classify_id" gorm:"type:bigint;"`                // 类型id
	CompanyId        int       `json:"company_id" gorm:"type:bigint;"`                 // 公司名
	TagId            int       `json:"tag_id" gorm:"type:bigint;"`                     // 标签id
	TypeId           int       `json:"type_id" gorm:"type:bigint;"`                    // 分类id
	CreateAt         time.Time `json:"create_at"`                                      // 创建时间
	DeleteAt         time.Time `json:"delete_at"`
	UpdateAt         time.Time `json:"update_at"`
	IsUse            int       `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}
