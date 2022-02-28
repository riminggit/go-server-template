package companyDelete

import (
	companyModel "go-server-template/model/company"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	companyHelper "go-server-template/src/company/helper"

	"github.com/gin-gonic/gin"
)

func DeleteService(c *gin.Context, params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}

	if params.RealDel == "1" {
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Delete(&companyModel.Company{}, "id IN ?", params.IDList).Error
			if delErr != nil {
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.NameList) > 0 {
			delErr := DB.DBLivingExample.Delete(&companyModel.Company{}, "company_name IN ?", params.NameList).Error
			if delErr != nil {
				res.Data = append(res.Data, "名称条件删除失败")
			}
		}
	} else {
		setData := map[string]interface{}{"is_use": 0, "DeleteAt": util.GetNowTimeUnix()}
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Model(&companyModel.Company{}).Where("id IN ?", params.IDList).Updates(setData).Error
			if delErr != nil {
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.NameList) > 0 {
			delErr := DB.DBLivingExample.Model(&companyModel.Company{}).Where("company_name IN ?", params.NameList).Updates(setData).Error
			if delErr != nil {
				res.Data = append(res.Data, "名称条件删除失败")
			}
		}
	}

	// 清除查询的redis缓存
	companyHelper.CleanRedisQuery()

	res.Code = e.SUCCESS
	return res
}
