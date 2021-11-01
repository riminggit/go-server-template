package wxDispose

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	projectConfig "go-server-template/config"
	"net/http"
)

// 这个函数以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXGetOpenIdService(code string) (*WXGetOpenIdResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, projectConfig.AppConfig.WXAPPConfig.APPID, projectConfig.AppConfig.WXAPPConfig.SECRET, code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXGetOpenIdResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}
	return &wxResp, nil
}

// CBC 模式
//解密
/**
* rawData 原始加密数据
* key  密钥
* iv  向量
 */
func WXDncryptService(param WXDncryptParams) (string, error) {

	data, err := base64.StdEncoding.DecodeString(param.Rawdata)
	key_b, err_1 := base64.StdEncoding.DecodeString(param.SessionKey)
	iv_b, _ := base64.StdEncoding.DecodeString(param.Iv)

	if err != nil {
		return "", err
	}
	if err_1 != nil {
		return "", err_1
	}

	dnData, err := AesCBCDncrypt(data, key_b, iv_b)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}
