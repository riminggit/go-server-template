package classifyUpdate

import (
	classifyModel "go-server-template/model/classify"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	classifyHelper "go-server-template/src/classify/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateClassifyService(c *gin.Context, params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}

	var resData []string
	for index, item := range params.Data {
		if item.ID != 0 {
			setData := classifyModel.Classify{
				UpdateAt: util.GetNowTimeUnix(),
			}
			if item.ClassifyName != "" {
				setData.ClassifyName = item.ClassifyName
			}
			if item.ImgUrl != "" {
				setData.ImgUrl = item.ImgUrl
			}
			if item.ImgSvg != "" {
				setData.ImgSvg = item.ImgSvg
			}
			if item.Rank != 0 {
				setData.Rank = item.Rank
			}
			res := DB.DBLivingExample.Model(&classifyModel.Classify{}).Where("id = ?", item.ID).Updates(setData)
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
	classifyHelper.CleanRedisQuery()

	res.Data = resData
	res.Code = e.SUCCESS
	return res
}
