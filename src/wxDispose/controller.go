package wxDispose

import (
	"encoding/json"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 微信获取openid和session_key,为了数据安全，不应该把session_key返回前端
// @Produce  json
// @Param code query string false "openid"
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
// @Param sessionKey query string false "sessionKey"
// @Param iv query string false "iv"
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

	// sessionkey不应该从客户端传过来，而是存在redis里比较好

}
