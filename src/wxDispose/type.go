package wxDispose

type WXGetOpenIdParams struct {
	Code string `json:"code"`
}

type WXDncryptParams struct {
	SessionKey string `json:"sessionKey"`
	Rawdata    string `json:"rawdata"`
	Iv         string `json:"iv"`
}

type WXGetOpenIdResp struct {
	OpenId string `json:"openid"`
	// SessionKey string `json:"sessionKey"`
	UnionId string `json:"unionid"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
