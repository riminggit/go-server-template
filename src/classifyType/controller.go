package classifyType

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	// "go-server-template/pkg/e"
	// "go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询二级分类
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query string false "id"
// @Param classify_name query string false "classify_name"
// @Param type_name query string false "type_name"
// @Param classify_id query string false "classify_id"
// @Param is_use query number false "is_use"
// @Router /api/type/query-type [get]
func QueryTypeController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonData := &QueryTypeParams{
		Id:           c.Query("id"),
		ClassifyName: c.Query("classify_name"),
		IsUse:        c.DefaultQuery("is_use", "1"),
		ClassifyId:   c.Query("classify_id"),
		TypeName:     c.Query("type_name"),
	}

	result := QueryTypeService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}
