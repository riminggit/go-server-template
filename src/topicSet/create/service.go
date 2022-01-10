package topicSetCreate

import (
	topicModel "go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context, params CreateParams) *CreateReturn {
	res := &CreateReturn{}
	res.Code = e.SUCCESS

	// 数组转字符串
	var topicList string
	if len(params.TopicSetIdList) > 0 {
		topicList = strings.Join(params.TopicSetIdList, ",")
	}

	createData := &topicModel.TopicSet{
		TopicSetIdList:     topicList,
		Name:               params.Name,
		TopicSetDifficulty: params.TopicSetDifficulty,
		TopicSetLevel:      params.TopicSetLevel,
		Remark:             params.Remark,
		CreateAt:           time.Now().Add(8 * time.Hour),
		TopicType:			params.TopicType,
		IsUse:              1,
	}

	err := DB.DBLivingExample.Model(&topicModel.TopicSet{}).Create(createData).Error
	if err != nil {
		res.Code = e.CREATE_DATA_FALE
		logging.Error(err)
		return res
	}
	return res
}
