package topicSetUpdate

import (
	"fmt"
	"go-server-template/pkg/app"
	"go-server-template/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
)

// @Summary 修改套题
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data query []st.UpdateParams false "data"
// @Router /api/topic/update-topic-set [post]
func UpdateController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &UpdateParams{}

	// 如果是里面包含数组得用这个解析
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	fmt.Println(jsonData,"jsonData")
	result := UpdateService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}

