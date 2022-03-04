package userTopicRead

import (
	"go-server-template/pkg/app"
	logging "go-server-template/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
)

// @Summary 新增阅题记录
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data query st.CreateParams false "data"
// @Router /api/user-topic/create-user-topic-read-record [post]]
func CreateController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &CreateParams{}
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}
	result := CreateUserTopicRead(c, *jsonData)
	appG.SimpleResponse(http.StatusOK, result.Code, result.Msg)
}
