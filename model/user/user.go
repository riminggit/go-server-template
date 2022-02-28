package userModel

import (
	"time"
)

// 文档https://gorm.io/zh_CN/docs/conventions.html

type User struct {
	ID        int       `json:"id" gorm:"type:bigint;column:id;comment:id;not null"` // id
	Email     string    `json:"email" gorm:"type:varchar(50);index;"`                // 唯一索引
	NickName  string    `json:"nick_name" gorm:"type:varchar(100);uniqueIndex;"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	Phone     string    `json:"phone" gorm:"type:varchar(13);index;"`
	AvatarUrl string    `json:"avatar_url" gorm:"type:varchar(255)"`
	ComeFrom  string    `json:"come_from" gorm:"type:varchar(20)"`
	Openid    string    `json:"openid" gorm:"type:varchar(50)"`
	Gender    int       `json:"gender" gorm:"type:int(3)"` // 1时是男性，值为2时是女性，值为0时是未知
	City      string    `json:"city" gorm:"type:varchar(50)"`
	Province  string    `json:"province" gorm:"type:varchar(50)"`
	Country   string    `json:"country" gorm:"type:varchar(50)"`
	Language  string    `json:"language" gorm:"type:varchar(30)"`
	Rawdata   string    `json:"rawdata" gorm:"type:varchar(255)"`
	Signature string    `json:"signature" gorm:"type:varchar(255)"`
	Iv        string    `json:"iv" gorm:"type:varchar(255)"`
	IsAdmin   int       `json:"is_admin" gorm:"type:int(3);default:0"`
	Birthday  time.Time `json:"birthday" gorm:"type:date;"`
	CreateAt  int64     `json:"create_at"`
	DeleteAt  int64     `json:"delete_at"`
	UpdateAt  int64     `json:"update_at"`
	IsUse     int       `json:"is_use" gorm:"type:int(3);default:1"`
}
