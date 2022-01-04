package classifyTypeUpdate

import (
	"github.com/gin-gonic/gin"
	"go-server-template/model/type"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/classifyType/helper"
	"strconv"
	"time"
)

func UpdateService(c *gin.Context, params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}

	var resData []string
	for index, item := range params.Data {
		if item.ID != 0 {
			setData := typeModel.Type{
				UpdateAt: time.Now().Add(8 * time.Hour),
			}
			if item.ClassifyId != 0 {
				setData.ClassifyId = item.ClassifyId
			}
			if item.TypeName != "" {
				setData.TypeName = item.TypeName
			}
			if item.Rank != 0 {
				setData.Rank = item.Rank
			}
			res := DB.DBLivingExample.Model(&typeModel.Type{}).Where("id = ?", item.ID).Updates(setData)
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
	classifyTypeHelper.CleanRedisQuery()
	res.Data = resData
	res.Code = e.SUCCESS
	return res
}
