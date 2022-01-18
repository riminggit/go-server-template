package userExperience

import (
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/app"
	"github.com/ugorji/go/codec"
	"go-server-template/pkg/log"
	"go-server-template/config"
	"go-server-template/pkg/e"
	"net/http"
)

// @Summary 管理员查询用户经验
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router /api/user/query-user-experience [post]
func QueryExperienceController(c *gin.Context) {
	appG := app.Gin{C: c}
	jsonString := app.GetPostJson(c)
	jsonData := &QueryParams{}

	// 如果是里面包含数组得用这个解析
	err := codec.NewDecoderBytes([]byte(jsonString), new(codec.JsonHandle)).Decode(jsonData)
	if err != nil {
		logging.Debug(err)
	}

	httpCode, errCode := app.JsonValid(c, jsonData)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	if jsonData.IsUse == "" {
		jsonData.IsUse = "1"
	}
	if jsonData.PageNum == 0 {
		jsonData.PageNum = 1
	}
	if jsonData.PageSize == 0 {
		jsonData.PageSize = projectConfig.AppConfig.BaseConfig.PAGE_SIZE
	}

	result := QueryUserExperience(c, *jsonData)
	appG.Response(http.StatusOK, result.Code, result.Data)
}


// @Summary 用户查询自身经验
// @Produce  json
// @Param Authorization	header string false "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router /api/user/query-user-experience-from-user [get]
func UserQueryExperienceController(c *gin.Context) {
	appG := app.Gin{C: c}

	result := UserQueryExperience(c)
	appG.Response(http.StatusOK, result.Code, result.Data)
}



