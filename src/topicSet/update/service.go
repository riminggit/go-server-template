package topicSetUpdate

import (
	topicModel "go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateService(c *gin.Context, params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}

	if params.ID != 0 {
		setData := topicModel.TopicSet{
			UpdateAt:           time.Now().Add(8 * time.Hour),
			TopicSetDifficulty: params.TopicSetDifficulty,
			TopicSetLevel:      params.TopicSetLevel,
			IsUse:              params.IsUse,
		}

		if len(params.TopicIdList) > 0 {
			topicList := strings.Join(params.TopicIdList, ",")
			setData.TopicIdList = topicList
		}
		if params.Name != "" {
			setData.Name = params.Name
		}
		if params.Remark != "" {
			setData.Remark = params.Remark
		}
		if params.TopicType != "" {
			setData.TopicType = params.TopicType
		}

		Res := DB.DBLivingExample.Model(&topicModel.TopicSet{}).Where("id = ?", params.ID).Updates(setData)
		if Res.Error != nil {
			res.Code = e.UPDATE_DATA_FALE
		}
	} else {
		res.Code = e.NO_DATA_EXISTS
	}

	res.Code = e.SUCCESS
	return res
}
