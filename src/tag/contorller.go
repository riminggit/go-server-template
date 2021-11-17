package tag

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	// "go-server-template/pkg/e"
	// "go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询标签
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query string false "id"
// @Param tag_name query string false "tag_name"
// @Param is_use query number false "is_use"
// @Router /api/tag/query-tag [get]
func QueryTagController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonData := &queryTagParams{
		Id:           c.Query("id"),
		TagName: c.Query("tag_name"),
		IsUse:        c.DefaultQuery("is_use", "1"),
	}

	result := queryTagService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}



