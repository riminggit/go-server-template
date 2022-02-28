package topicModel

type TopicSet struct {
	ID                 int    `json:"id" gorm:"type:bigint;column:id;comment:id;not null"` // id
	Name               string `json:"name" gorm:"type:varchar(50);"`                       // 套题名称
	TopicIdList        string `json:"topic_id_list" gorm:"type:text;"`                     // 题目id列表
	TopicNum           int    `json:"topic_num" gorm:"type:int(5);"`                       // 题目数量
	TopicSetDifficulty int    `json:"topic_set_difficulty" gorm:"type:int(3);"`            // 题目难度，由所有的题目难度指数相加后除题目数
	TopicSetLevel      int    `json:"topic_set_level" gorm:"type:int(3);"`                 // 题目等级，由所有的题目等级指数相加后除题目数
	TopicType          string `json:"topic_type" gorm:"type:int(3);"`                      // 练习：1  考试 ：2
	Remark             string `json:"remark" gorm:"type:varchar(255);"`                    // 备注
	CreateAt           int64  `json:"create_at"`                                           // 创建时间
	DeleteAt           int64  `json:"delete_at"`
	UpdateAt           int64  `json:"update_at"`
	IsUse              int    `json:"is_use" gorm:"type:int(3);not null;default:1"` // 是否删除:0 删除 1未删除
}
