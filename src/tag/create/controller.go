package tagCreate

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	"github.com/ugorji/go/codec"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 新增标签
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data query []st.CreateParams false "data"
// @Router /api/tag/create-tag-multiple [post]
func CreateController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &CreateParams{}

	// 如果是里面包含数组得用这个解析
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	result := CreateService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}

