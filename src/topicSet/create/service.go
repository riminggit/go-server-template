package topicSetCreate

import (
	"github.com/gin-gonic/gin"
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	"go-server-template/pkg/snowflake"
	"strings"
	"time"
)

func CreateService(c *gin.Context, params CreateParams) *CreateReturn {
	res := &CreateReturn{}
	res.Code = e.SUCCESS

	manageResp := TopSetDifficultyAndLevelCompute(c, params.TopicSetIdList)

	if manageResp.Code != e.SUCCESS {
		res.Code = e.CREATE_DATA_FALE
		return res
	}

	// 字符串分割为数组 strings.Split(s, sep string) []string

	// 数组转字符串
	var topicList string
	if len(params.TopicSetIdList) > 0 {
		// 数组转字符串
		topicList = strings.Join(params.TopicSetIdList, ",")
	}

	createData := &topicModel.TopicSet{
		ID:                 snowflake.GenerateID(1),
		TopicSetIdList:     topicList,
		Name:               params.Name,
		TopicSetDifficulty: manageResp.Data.Difficulty,
		TopicSetLevel:      manageResp.Data.Level,
		Remark:             params.Remark,
		CreateAt:           time.Now().Add(8 * time.Hour),
		TopicType:          params.TopicType,
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
