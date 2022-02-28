package userModel

// 用redis存，然后查的时候直接查redis，同步更新到这个表里面，如果有数据就是已读，没有就是未读
type UserTopicRead struct {
	ID       int   `json:"id" gorm:"column:id;AUTO_INCREMENT;comment:id;not null"` // id
	UserId   int   `json:"user_id" gorm:"type:bigint;not null"`                    // 用户id
	TopicId  int   `json:"topic_id" gorm:"type:bigint;not null"`                   // 对应题目id
	IsRead   int   `json:"is_read" gorm:"type:int(3);"`                            // 是否阅读 0 未读  1已读
	ReadNum  int   `json:"read_num" gorm:"type:int(4);"`                           // 阅读次数
	CreateAt int64 `json:"create_at"`                                              // 创建时间
	UpdateAt int64 `json:"update_at"`
}
