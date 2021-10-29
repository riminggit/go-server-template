package userLogin

import (
	userModel "go-server-template/model/user"
)

// tag想多加一个类型空格就可以了，例如 `json:"openid" valid:"Alpha"`
type WXUserParams struct {
	WXUserCreateParams
	SessionKey string `json:"sessionKey"`
}

type WXUserCreateParams struct {
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
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
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
