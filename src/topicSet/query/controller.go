package topicSetQuery


import (
	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询用户答题记录
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query []string false "id"
// @Param name query []string false "name"
// @Param topic_set_id_list query string false "topic_set_id_list"
// @Param topic_set_difficulty query string false "topic_set_difficulty"
// @Param topic_set_level query string false "topic_set_level"
// @Param create_at query []string false "create_at"
// @Param delete_at query []string false "delete_at"
// @Param update_at query []string false "update_at"
// @Param is_use query string false "is_use"
// @Param pageNum query int false "pageNum"
// @Param pageSize query int false "pageSize"
// @Router /api/topic/query-topic-set [post]
func QueryTopicSetController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonString := app.GetPostJson(c)
	jsonData := &QueryTopicSetParams{}

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

	result := QueryTopicSetService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}
