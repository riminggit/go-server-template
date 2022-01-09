package feedback

import (
	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
	"go-server-template/config"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询用户答题记录
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query []string false "id"
// @Param user_id query []string false "user_id"
// @Param topic_id query int false "topic_id"
// @Param feedback_type query int false "feedback_type"
// @Param grade query int false "grade"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Param pageNum query int false "pageNum"
// @Param pageSize query int false "pageSize"
// @Router /api/feedback/query-feedback [post]
func QueryController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &QueryParams{}

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

	if jsonData.PageNum == 0 {
		jsonData.PageNum = 1
	}
	if jsonData.PageSize == 0 {
		jsonData.PageSize = projectConfig.AppConfig.BaseConfig.PAGE_SIZE
	}
	if jsonData.IsUse == "" {
		jsonData.IsUse = "1"
	}

	result := QueryService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}



// @Summary 用户查询答题记录
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query []string false "id"
// @Param topic_id query int false "topic_id"
// @Param feedback_type query int false "feedback_type"
// @Param grade query int false "grade"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Router /api/feedback/query-feedback-user [post]
func UserQueryController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &UserQueryParams{}

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

	result := UserQueryService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}
