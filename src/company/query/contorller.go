package companyQuery

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	"github.com/ugorji/go/codec"
	// "go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"net/http"
)

// @Summary 查询公司
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query string false "id"
// @Param company_name query string false "company_name"
// @Param is_use query number false "is_use"
// @Router /api/company/query-company [get]
func QueryCompanyController(c *gin.Context) {
	appG := app.Gin{C: c}

	jsonData := &queryCompanyParams{
		Id:           c.Query("id"),
		CompanyName: c.Query("company_name"),
		IsUse:        c.DefaultQuery("is_use", "1"),
	}

	result := queryCompanyService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}

// @Summary 查询公司（多个参数）
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param id query []string false "id"
// @Param company_name query []string false "company_name"
// @Param is_use query string false "is_use"
// @Router /api/company/query-company [post]
func QueryCompanyMultipleController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &queryCompanyMultipleParams{}

	// 如果是里面包含数组得用这个解析
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	if jsonData.IsUse == "" {
		jsonData.IsUse = "1"
	}

	result := queryCompanyMultipleService(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}




