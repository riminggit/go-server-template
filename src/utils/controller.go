package utils

import (
	"go-server-template/pkg/app"
	logging "go-server-template/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ugorji/go/codec"
)

// @Summary 添加颜色
// @Produce  json
// @Param code query string false "openid"
// @Router /api/utils/add-color [post]
func AddColor(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &AddParams{}
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	result := AddColorSerices(c, *jsonData)
	appG.SimpleResponse(http.StatusOK, result.Code, result.Msg)
}

// @Summary 查询颜色
// @Produce  json
// @Param code query string false "openid"
// @Router /api/utils/query-color [get]
func QueryColor(c *gin.Context) {
	appG := app.Gin{C: c}

	result := QueryColorService(c)
	appG.Response(http.StatusOK, result.Code, result.Data)
}
