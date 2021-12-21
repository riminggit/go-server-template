package midTopicTagQuery

import (
	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
	// "go-server-template/config"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询topic-tag-关系
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query []string false "id"
// @Param topic_id query string false "topic_id"
// @Param topic_id query int false "topic_id"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Router /api/topic/query-topic/tag/mid [post]
func QueryTopicTagMidController(c *gin.Context) {
	// 做一下判断不允许传参过多
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &QueryTopicTagMidParams{}

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

	result := QueryTopicTagMid(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}
