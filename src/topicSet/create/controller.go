package topicSetCreate

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	"github.com/ugorji/go/codec"
	"go-server-template/pkg/log"
	"net/http"
)


// @Summary 新增题目
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data query st.CreateParams false "data"
// @Router /api/topic/create-topic-set [post]]
func CreateController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &CreateParams{}
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}
	result := CreateService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}

