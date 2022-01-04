package companyDelete

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	"github.com/ugorji/go/codec"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 删除公司
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data query []st.DeleteParams false "data"
// @Router /api/company/delete-company-multiple [post]]
func DeleteController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &DeleteParams{}

	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	result := DeleteService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}

