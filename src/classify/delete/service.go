package classifyDelete

import (
	classifyModel "go-server-template/model/classify"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	classifyHelper "go-server-template/src/classify/helper"

	"github.com/gin-gonic/gin"
)

func DeleteClassifyService(c *gin.Context, params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}

	if params.RealDel == "1" {
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Delete(&classifyModel.Classify{}, "id IN ?", params.IDList).Error
			if delErr != nil {
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.ClassifyNameList) > 0 {
			delErr := DB.DBLivingExample.Delete(&classifyModel.Classify{}, "classify_name IN ?", params.ClassifyNameList).Error
			if delErr != nil {
				res.Data = append(res.Data, "分类名条件删除失败")
			}
		}
	} else {
		setData := map[string]interface{}{"is_use": 0, "DeleteAt": util.GetNowTimeUnix()}
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Model(&classifyModel.Classify{}).Where("id IN ?", params.IDList).Updates(setData).Error
			if delErr != nil {
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.ClassifyNameList) > 0 {
			delErr := DB.DBLivingExample.Model(&classifyModel.Classify{}).Where("classify_name IN ?", params.ClassifyNameList).Updates(setData).Error
			if delErr != nil {
				res.Data = append(res.Data, "分类名条件删除失败")
			}
		}
	}

	// 清除查询的redis缓存
	classifyHelper.CleanRedisQuery()

	res.Code = e.SUCCESS
	return res
}
