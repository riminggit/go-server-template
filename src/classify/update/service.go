package classifyUpdate

import (
	"go-server-template/model/classify"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateClassifyService(c *gin.Context, params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}

	var resData []string
	for index, item := range params.Data {
		if item.ID != 0 {
			setData := classifyModel.Classify{
				ClassifyName: item.ClassifyName,
				ImgUrl:       item.ImgUrl,
				ImgSvg:       item.ImgSvg,
				Rank:         item.Rank,
				UpdateAt:     time.Now().Add(8 * time.Hour),
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
	res.Data = resData
	res.Code = e.SUCCESS
	return res
}

