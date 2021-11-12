package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
	"go-server-template/config"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询用户信息
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param phone query string false "手机号"
// @Param email query string false "email"
// @Param nick_name query string false "nick_name"
// @Param come_from query string false "come_from"
// @Param gender query string false "gender"
// @Param city query []string false "city"
// @Param province query []string false "province"
// @Param country query []string false "country"
// @Param language query []string false "language"
// @Param is_admin query string false "is_admin"
// @Param birthday query []string false "birthday"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Param pageNum query int false "pageNum"
// @Param pageSize query int false "pageSize"
// @Success 200 {object} Response
// @Router /api/user/account-manage/binding-phone [post]
func QueryUserDataController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &QueryUserParams{}

	// 如果是里面包含数组得用这个解析
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	httpCode, errCode := app.JsonValid(c, jsonData)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	if jsonData.IsUse == "" {
		jsonData.IsUse = "1"
	}
	if jsonData.PageNum == 0 {
		jsonData.PageNum = 1
	}
	if jsonData.PageSize == 0 {
		jsonData.PageSize = projectConfig.AppConfig.BaseConfig.PAGE_SIZE
	}

	result := QueryUserDataService(c, *jsonData)

	appG.Response(http.StatusOK, result.Code, result.Data)
}
