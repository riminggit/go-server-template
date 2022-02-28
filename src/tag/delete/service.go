package tagDelete

import (
	tagModel "go-server-template/model/tag"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	tagHelper "go-server-template/src/tag/helper"

	"github.com/gin-gonic/gin"
)

func DeleteService(c *gin.Context, params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}

	if params.RealDel == "1" {
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Delete(&tagModel.Tag{}, "id IN ?", params.IDList).Error
			if delErr != nil {
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.NameList) > 0 {
			delErr := DB.DBLivingExample.Delete(&tagModel.Tag{}, "tag_name IN ?", params.NameList).Error
			if delErr != nil {
				res.Data = append(res.Data, "名称条件删除失败")
			}
		}
	} else {
		setData := map[string]interface{}{"is_use": 0, "DeleteAt": util.GetNowTimeUnix()}
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Model(&tagModel.Tag{}).Where("id IN ?", params.IDList).Updates(setData).Error
			if delErr != nil {
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.NameList) > 0 {
			delErr := DB.DBLivingExample.Model(&tagModel.Tag{}).Where("tag_name IN ?", params.NameList).Updates(setData).Error
			if delErr != nil {
				res.Data = append(res.Data, "名称条件删除失败")
			}
		}
	}

	// 清除查询的redis缓存
	tagHelper.CleanRedisQuery()

	res.Code = e.SUCCESS
	return res
}
