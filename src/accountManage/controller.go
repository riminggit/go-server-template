package account

import (
	"encoding/json"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 绑定/修改手机号
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param phone query string false "手机号"
// @Router /api/user/account-manage/binding-phone [post]
func ChangePhoneController(c *gin.Context) {
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
	result := ChangePhoneService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, nil)
}
