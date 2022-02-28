package userloginAndExit

import (
	userModel "go-server-template/model/user"
)

// tag想多加一个类型空格就可以了，例如 `json:"openid" valid:"Alpha"`
type WXUserCreateParams struct {
	ID        int    `json:"id"` // id
	Openid    string `json:"openid"`
	ComeFrom  string `json:"come_from"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Language  string `json:"language"`
	Rawdata   string `json:"rawdata"`
	Signature string `json:"signature"`
	Iv        string `json:"iv"`
	Phone     string `json:"phone"`
	UpdateAt  int64  `json:"update_at"`
	CreateAt  int64  `json:"create_at"`
}

type LoginParams struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	NickName  string `json:"nick_name"`
	Password  string `json:"password" valid:"Required;"`
	ComeFrom  string `json:"come_from"`
	LoginType string `json:"login_type"`
}

type LoginReturnData struct {
	Token    string         `json:"token"`
	Code     int            `json:"code"`
	UserInfo userModel.User `json:"userInfo"`
	Msg      string         `json:"msg"`
}

type WXLoginResultData struct {
	Token string `json:"token"`
}

type LayoutReturnData struct {
	Code int `json:"code"`
}
