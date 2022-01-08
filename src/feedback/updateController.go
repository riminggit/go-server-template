package feedback

import (
	"fmt"
	"go-server-template/pkg/app"
	"go-server-template/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
)

// @Summary 回复反馈
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data query []st.AnswerFeedbackParams false "data"
// @Router /api/feedback/answer-feedback [post]
func AnswerFeedbackController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &AnswerFeedbackParams{}

	// 如果是里面包含数组得用这个解析
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	fmt.Println(jsonData,"jsonData")
	result := AnswerFeedbackService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}


// @Summary 用户修改反馈
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data query []st.UserUpdateParams false "data"
// @Router /api/feedback/update-feedback-user [post]
func UserUpdateController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &UserUpdateParams{}

	// 如果是里面包含数组得用这个解析
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	fmt.Println(jsonData,"jsonData")
	result := UserUpdateService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}

