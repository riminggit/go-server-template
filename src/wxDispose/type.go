package wxDispose

type WXGetOpenIdParams struct {
	Code string `json:"code"`
}

type WXDncryptParams struct {
	SessionKey string `json:"session_key"`
	RawData    string `json:"rawData"`
	Iv         string `json:"iv"`
}

type WXGetOpenIdResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
