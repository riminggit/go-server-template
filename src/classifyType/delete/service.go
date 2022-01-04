package classifyTypeDelete

import (
	"github.com/gin-gonic/gin"
	"go-server-template/model/type"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/classifyType/helper"
	"time"
)

func DeleteService(c *gin.Context, params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}

	if params.RealDel == "1" {
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Delete(&typeModel.Type{}, "id IN ?", params.IDList).Error
			if delErr != nil {
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.NameList) > 0 {
			delErr := DB.DBLivingExample.Delete(&typeModel.Type{}, "type_name IN ?", params.NameList).Error
			if delErr != nil {
				res.Data = append(res.Data, "分类名条件删除失败")
			}
		}
	} else {
		setData := map[string]interface{}{"is_use": 0, "DeleteAt": time.Now().Add(8 * time.Hour)}
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Model(&typeModel.Type{}).Where("id IN ?", params.IDList).Updates(setData).Error
			if delErr != nil {
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.NameList) > 0 {
			delErr := DB.DBLivingExample.Model(&typeModel.Type{}).Where("type_name IN ?", params.NameList).Updates(setData).Error
			if delErr != nil {
				res.Data = append(res.Data, "分类名条件删除失败")
			}
		}
	}

	// 清除查询的redis缓存
	classifyTypeHelper.CleanRedisQuery()

	res.Code = e.SUCCESS
	return res
}
