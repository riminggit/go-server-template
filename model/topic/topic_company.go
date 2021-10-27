package topicModel

type TopicCompany struct {
	ID        int    `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	TopicId   int    `gorm:"type:int(11);not null"`                                  // 对应题目id
	CompanyId int    `gorm:"type:int(11);not null"`                                  // 公司名
	TopicTime string `gorm:"type:varchar(30)"`                                       // 题目时间
	IsUse     int    `gorm:"type:int(3);not null;default:1"`                         // 是否删除:0 删除 1未删除
}
