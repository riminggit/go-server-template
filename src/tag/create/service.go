package tagCreate

import (
	tagModel "go-server-template/model/tag"
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
	createData := []tagModel.Tag{}
	for index, item := range params.Data {
		var queryInfo []tagModel.Tag
		DB.DBLivingExample.Where("is_use = ?", 0).Where("tag_name = ?", item.TagName).Model(&tagModel.Tag{}).Find(&queryInfo)
		if len(queryInfo) > 0 {
			hasOld = true
			updateData := tagModel.Tag{
				TagName:  item.TagName,
				UpdateAt: util.GetNowTimeUnix(),
				IsUse:    1,
			}
			updateErr := DB.DBLivingExample.Model(&tagModel.Tag{}).Where("id = ?", queryInfo[0].ID).Updates(updateData).Error
			if updateErr != nil {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据更新出错"
				resData = append(resData, msg)
			} else {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据恢复更新成功"
				resData = append(resData, msg)
			}
		} else {
			setData := tagModel.Tag{
				ID:       snowflake.GenerateID(1),
				TagName:  item.TagName,
				CreateAt: util.GetNowTimeUnix(),
				IsUse:    1,
			}
			createData = append(createData, setData)
		}
	}

	res.Data = resData

	err := DB.DBLivingExample.Model(&tagModel.Tag{}).Create(createData).Error
	if err != nil {
		if !hasOld {
			res.Code = e.CREATE_DATA_FALE
		}
		return res
	}

	res.Code = e.SUCCESS
	return res
}
