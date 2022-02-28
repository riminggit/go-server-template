package topicModel

type TopicCompany struct {
	ID        int    `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	TopicId   int    `gorm:"type:bigint;not null"`                                   // 对应题目id
	CompanyId int    `gorm:"type:bigint;not null"`                                   // 公司名
	TopicTime string `gorm:"type:varchar(30)"`                                       // 题目时间
	CreateAt  int64  `json:"create_at"`                                              // 创建时间
	DeleteAt  int64  `json:"delete_at"`
	UpdateAt  int64  `json:"update_at"`
}
