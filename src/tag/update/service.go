package tagUpdate

import (
	tagModel "go-server-template/model/tag"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	tagHelper "go-server-template/src/tag/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateService(c *gin.Context, params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}

	var resData []string
	for index, item := range params.Data {
		if item.ID != 0 {
			setData := tagModel.Tag{
				UpdateAt: util.GetNowTimeUnix(),
			}
			if item.TagName != "" {
				setData.TagName = item.TagName
			}
			res := DB.DBLivingExample.Model(&tagModel.Tag{}).Where("id = ?", item.ID).Updates(setData)
			if res.Error != nil {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据更新出错"
				resData = append(resData, msg)
			} else {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据更新成功"
				resData = append(resData, msg)
			}
		} else {
			msg := "第" + strconv.Itoa((index + 1)) + "条数据缺少id"
			resData = append(resData, msg)
		}

	}

	// 清除查询的redis缓存
	tagHelper.CleanRedisQuery()

	res.Data = resData
	res.Code = e.SUCCESS
	return res
}
