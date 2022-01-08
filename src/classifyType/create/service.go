package classifyTypeCreate

import (
	typeModel "go-server-template/model/type"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context, params CreateParams) *CreateReturn {
	res := &CreateReturn{}

	var resData []string
	var hasOld bool
	createData := []typeModel.Type{}
	for index, item := range params.Data {
		var queryInfo []typeModel.Type
		DB.DBLivingExample.Where("is_use = ?", 0).Where("type_name = ?", item.TypeName).Model(&typeModel.Type{}).Find(&queryInfo)
		if len(queryInfo) > 0 {
			hasOld = true
			updateData := typeModel.Type{
				ClassifyId: item.ClassifyId,
				Rank:       item.Rank,
				UpdateAt:   time.Now().Add(8 * time.Hour),
				IsUse:      1,
			}
			updateErr := DB.DBLivingExample.Model(&typeModel.Type{}).Where("id = ?", queryInfo[0].ID).Updates(updateData).Error
			if updateErr != nil {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据更新出错"
				resData = append(resData, msg)
			} else {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据恢复更新成功"
				resData = append(resData, msg)
			}
		} else {
			setData := typeModel.Type{
				ClassifyId: item.ClassifyId,
				TypeName:   item.TypeName,
				Rank:       item.Rank,
				CreateAt:   time.Now().Add(8 * time.Hour),
			}
			createData = append(createData, setData)
		}
	}

	res.Data = resData

	err := DB.DBLivingExample.Model(&typeModel.Type{}).Create(createData).Error
	if err != nil {
		if !hasOld {
			res.Code = e.CREATE_DATA_FALE
		}
		return res
	}

	res.Code = e.SUCCESS
	return res
}
