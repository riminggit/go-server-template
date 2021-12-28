package classifyUpdate

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	"github.com/ugorji/go/codec"
	// "go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 修改分类- 校验管理员
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data query []st.ClassifyParams false "data"
// @Router /api/classify/update-classify-multiple [post]]
func UpdateClassifyController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &UpdateParams{}

	// 如果是里面包含数组得用这个解析
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	result := UpdateClassifyService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}

