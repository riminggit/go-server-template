package classifyCreate

import (
	classifyModel "go-server-template/model/classify"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/pkg/snowflake"
	util "go-server-template/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateClassifyService(c *gin.Context, params CreateParams) *CreateReturn {
	res := &CreateReturn{}

	var resData []string
	var hasOld bool
	createData := []classifyModel.Classify{}
	for index, item := range params.Data {
		var classifyInfo []classifyModel.Classify
		DB.DBLivingExample.Where("is_use = ?", 0).Where("classify_name = ?", item.ClassifyName).Model(&classifyModel.Classify{}).Find(&classifyInfo)
		if len(classifyInfo) > 0 {
			hasOld = true
			updateData := classifyModel.Classify{
				ImgUrl:   item.ImgUrl,
				ImgSvg:   item.ImgSvg,
				Rank:     item.Rank,
				UpdateAt: util.GetNowTimeUnix(),
				IsUse:    1,
			}
			updateErr := DB.DBLivingExample.Model(&classifyModel.Classify{}).Where("id = ?", classifyInfo[0].ID).Updates(updateData).Error
			if updateErr != nil {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据更新出错"
				resData = append(resData, msg)
			} else {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据恢复更新成功"
				resData = append(resData, msg)
			}
		} else {
			setData := classifyModel.Classify{
				ID:           snowflake.GenerateID(1),
				ClassifyName: item.ClassifyName,
				ImgUrl:       item.ImgUrl,
				ImgSvg:       item.ImgSvg,
				Rank:         item.Rank,
				CreateAt:     util.GetNowTimeUnix(),
			}
			createData = append(createData, setData)
		}
	}

	res.Data = resData

	err := DB.DBLivingExample.Model(&classifyModel.Classify{}).Create(createData).Error
	if err != nil {
		if !hasOld {
			res.Code = e.CREATE_DATA_FALE
		}
		return res
	}

	res.Code = e.SUCCESS
	return res
}
