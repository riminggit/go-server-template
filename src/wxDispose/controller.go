package wxDispose

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 微信获取openid和session_key
// @Produce  json
// @Param code query string false "openid"
// @Success 200 {object} Response
// @Router /api/user-login/wxapp-get-openid [post]
func WXGetOpenIdController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &WXGetOpenIdParams{}
	err := json.Unmarshal([]byte(jsonString), jsonData)
	if err != nil {
		logging.Debug(err)
	}

	resCode := e.SUCCESS
	result, err := WXGetOpenIdService(jsonData.Code)
	if err != nil {
		resCode = e.ERROR
	}
	appG.Response(http.StatusOK, resCode, result)
}

// @Summary 微信解密
// @Produce  json
// @Param rawData query string false "rawData"
// @Param session_key query string false "session_key"
// @Param iv query string false "iv"
// @Success 200 {object} Response
// @Router /api/user-login/wxapp-dncrypt [post]
func WXDncryptController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &WXDncryptParams{}
	err := json.Unmarshal([]byte(jsonString), jsonData)
	if err != nil {
		logging.Debug(err)
	}

	resCode := e.SUCCESS
	result, err := WXDncryptService(*jsonData)
	if err != nil {
		resCode = e.ERROR
	}
	appG.Response(http.StatusOK, resCode, result)

	// 解密后数据
	// {
	// 	"phoneNumber": "13580006666",
	// 	"purePhoneNumber": "13580006666",
	// 	"countryCode": "86",
	// 	"watermark":
	// 	{
	// 		"appid":"APPID",
	// 		"timestamp": TIMESTAMP
	// 	}
	// }
}
