package classifyDelete

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
// @Router /api/classify/delete-classify-multiple [post]]
func DeleteClassifyController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &DeleteParams{}

	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	result := DeleteClassifyService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}

