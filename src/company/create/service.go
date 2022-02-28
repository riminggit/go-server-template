package companyCreate

import (
	companyModel "go-server-template/model/company"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/pkg/snowflake"
	util "go-server-template/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context, params CreateParams) *CreateReturn {
	res := &CreateReturn{}

	var resData []string
	var hasOld bool
	createData := []companyModel.Company{}
	for index, item := range params.Data {
		var queryInfo []companyModel.Company
		DB.DBLivingExample.Where("is_use = ?", 0).Where("company_name = ?", item.CompanyName).Model(&companyModel.Company{}).Find(&queryInfo)
		if len(queryInfo) > 0 {
			hasOld = true
			updateData := companyModel.Company{
				CompanyName: item.CompanyName,
				ImgUrl:      item.ImgUrl,
				ImgSvg:      item.ImgSvg,
				UpdateAt:    util.GetNowTimeUnix(),
				IsUse:       1,
			}
			updateErr := DB.DBLivingExample.Model(&companyModel.Company{}).Where("id = ?", queryInfo[0].ID).Updates(updateData).Error
			if updateErr != nil {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据更新出错"
				resData = append(resData, msg)
			} else {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据恢复更新成功"
				resData = append(resData, msg)
			}
		} else {
			setData := companyModel.Company{
				ID:          snowflake.GenerateID(1),
				CompanyName: item.CompanyName,
				ImgUrl:      item.ImgUrl,
				ImgSvg:      item.ImgSvg,
				CreateAt:    util.GetNowTimeUnix(),
				IsUse:       1,
			}
			createData = append(createData, setData)
		}
	}

	res.Data = resData

	err := DB.DBLivingExample.Model(&companyModel.Company{}).Create(createData).Error
	if err != nil {
		if !hasOld {
			res.Code = e.CREATE_DATA_FALE
		}
		return res
	}

	res.Code = e.SUCCESS
	return res
}
