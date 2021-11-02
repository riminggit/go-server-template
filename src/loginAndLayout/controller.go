package userLoginAndLayout

import (
	"encoding/json"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 微信小程序登陆相关方法
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param openid query string false "openid"
// @Param come_from query string false "微信"
// @Param nick_name query string false "用户名"
// @Param avatar_url query string false "用户头像"
// @Param gender query int false "性别"
// @Param city query string false "城市"
// @Param province query string false "省份"
// @Param country query string false "国家"
// @Param language query string false "语言"
// @Param rawdata query string false "rawdata"
// @Param signature query string false "signature"
// @Param encrypteddata query string false "encrypteddata"
// @Param iv query string false "iv"
// @Success 200 {object} Response
// @Router /api/user-login/wxLogin [post]
func UserWXLoginController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &WXUserCreateParams{}
	err := json.Unmarshal([]byte(jsonString), jsonData)
	if err != nil {
		logging.Debug(err)
	}

	httpCode, errCode := app.JsonValid(c, jsonData)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	result := UserWXLoginService(c, *jsonData)

	var returnData interface{}
	resData := WXLoginResultData{}
	resData.Token = result.Token
	if result.Code == e.SUCCESS {
		returnData = resData
	}

	appG.Response(http.StatusOK, result.Code, returnData)
}

// @Summary 登陆相关方法,分开两个方法是为了避免以后或许想对应不同端登陆可以做一些不同的操作
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param come_from query string false "登陆来源"
// @Param login_type query string false "登陆方式"
// @Param nick_name query string false "用户名"
// @Param phone query string false "手机号"
// @Param email query string false "邮箱"
// @Param password query string false "密码"
// @Success 200 {object} Response
// @Router /api/user-login/wxLogin [post]
func LoginController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &LoginParams{}
	err := json.Unmarshal([]byte(jsonString), jsonData)
	if err != nil {
		logging.Debug(err)
	}

	httpCode, errCode := app.JsonValid(c, jsonData)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	result := LoginService(c, *jsonData)

	var returnData interface{}
	resData := WXLoginResultData{}
	resData.Token = result.Token
	if result.Code == e.SUCCESS {
		returnData = resData
	}

	appG.Response(http.StatusOK, result.Code, returnData)
}
