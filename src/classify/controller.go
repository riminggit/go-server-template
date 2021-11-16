package classify

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	// "go-server-template/pkg/e"
	// "go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询分类
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query string false "id"
// @Param classify_name query string false "classify_name"
// @Param rank query string false "rank"
// @Param is_use query number false "is_use"
// @Success 200 {object} Response
// @Router /api/classify/query-classify [get]
func QueryClassifyController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonData := &queryClassifyParams{
		Id:           c.Query("id"),
		ClassifyName: c.Query("classify_name"),
		IsUse:        c.DefaultQuery("is_use", "1"),
		Rank:         c.Query("rank"),
	}

	result := queryClassifyService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}



// @Summary 查询分类以及对应二级分类
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query string false "id"
// @Param classify_name query string false "classify_name"
// @Param rank query string false "rank"
// @Param is_use query number false "is_use"
// @Success 200 {object} Response
// @Router /api/classify/query-classify-type [get]
func QueryClassifyAndTypeController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonData := &queryClassifyParams{
		Id:           c.Query("id"),
		ClassifyName: c.Query("classify_name"),
		IsUse:        c.DefaultQuery("is_use", "1"),
		Rank:         c.Query("rank"),
	}

	result := queryClassifyAndTypeService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}
