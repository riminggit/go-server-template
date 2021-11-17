package company

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	// "go-server-template/pkg/e"
	// "go-server-template/pkg/log"
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



