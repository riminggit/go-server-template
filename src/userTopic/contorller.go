package userTopic

import (
	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
	"go-server-template/config"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询用户新增题目
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query string false "id"
// @Param tag_name query string false "tag_name"
// @Param is_use query number false "is_use"
// @Router /api/user-topic/query-user-topic [post]
func QueryUserTopicController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &queryUserTopicParams{}

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

	result := queryUserTopicService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}
