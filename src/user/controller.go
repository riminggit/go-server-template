package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询用户信息
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param phone query string false "手机号"
// @Success 200 {object} Response
// @Router /api/user/account-manage/binding-phone [get]
func QueryUserInfoController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &ChangePhoneParams{}
	err := json.Unmarshal([]byte(jsonString), jsonData)
	if err != nil {
		logging.Debug(err)
	}

	httpCode, errCode := app.JsonValid(c, jsonData)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	result := QueryUserInfoService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, nil)
}
