package topicQuery

import (
	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
	"go-server-template/config"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询题目
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query []string false "id"
// @Param title query string false "title"
// @Param question_type query []string false "question_type"
// @Param degree query []string false "degree"
// @Param level query []string false "level"
// @Param is_base_topic query string false "is_base_topic"
// @Param is_important_topic query string false "is_important_topic"
// @Param relation query []string false "relation"
// @Param come_from query []string false "come_from"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Param pageNum query int false "pageNum"
// @Param pageSize query int false "pageSize"
// @Router /api/topic/query-topic [post]
func QueryTopicController(c *gin.Context) {
	// 做一下判断不允许传参过多
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &QueryTopicParams{}

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

	if len(jsonData.Relation) > 0 {
		result := QueryTopicRelationService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	} else {
		result := QueryTopicService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	}
}

// @Summary 查询题目-classify
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param classify_id query []string false "classify_id"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Param pageNum query int false "pageNum"
// @Param pageSize query int false "pageSize"
// @Router /api/topic/query-topic/classify [post]
func QueryTopicClassifyController(c *gin.Context) {
	// 做一下判断不允许传参过多
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &queryTopicFromClassifyParams{}

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

	if len(jsonData.Relation) > 0 {
		result := QueryTopicFromClassifyRelationService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	} else {
		result := QueryTopicFromClassifyService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	}

}

// @Summary 查询题目-company
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param company_id query []string false "company_id"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Param pageNum query int false "pageNum"
// @Param pageSize query int false "pageSize"
// @Router /api/topic/query-topic/company [post]
func QueryTopicCompanyController(c *gin.Context) {
	// 做一下判断不允许传参过多
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &queryTopicFromCompanyParams{}

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

	if len(jsonData.Relation) > 0 {
		result := QueryTopicFromCompanyRelationService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	} else {
		result := QueryTopicFromCompanyService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	}

}

// @Summary 查询题目-tag
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param tag_id query []string false "tag_id"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Param pageNum query int false "pageNum"
// @Param pageSize query int false "pageSize"
// @Router /api/topic/query-topic/tag [post]
func QueryTopicTagController(c *gin.Context) {
	// 做一下判断不允许传参过多
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &queryTopicFromTagParams{}

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

	if len(jsonData.Relation) > 0 {
		result := QueryTopicFromTagRelationService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	} else {
		result := QueryTopicFromTagService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	}
}

// @Summary 查询题目-type
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param type_id query []string false "type_id"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Param pageNum query int false "pageNum"
// @Param pageSize query int false "pageSize"
// @Router /api/topic/query-topic/type [post]
func QueryTopicTypeController(c *gin.Context) {
	// 做一下判断不允许传参过多
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &queryTopicFromTypeParams{}

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

	if len(jsonData.Relation) > 0 {
		result := QueryTopicFromTypeRelationService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	} else {
		result := QueryTopicFromTypeService(c, *jsonData)
		appG.Response(http.StatusOK, result.Code, result.Data)
	}
}
