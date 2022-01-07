package midTopicClassifyUpdate

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicClassify/create"
	"go-server-template/src/midTopicClassify/helper"
	"time"
	"gorm.io/gorm"
)

func UpdateService(params UpdateParams, db *gorm.DB) error {
	var queryInfo []topicModel.TopicClassify
	queryFun := db
	queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	queryFun = queryFun.Where("classify_id = ?", params.ClassifyId)
	queryFun.Model(&topicModel.TopicClassify{}).Find(&queryInfo)

	if len(queryInfo) > 0 {
		setData := topicModel.TopicClassify{
			UpdateAt:   time.Now().Add(8 * time.Hour),
			ClassifyId: params.NewClassifyId,
		}
		res := queryFun.Updates(setData)
		// 清除查询的redis缓存
		thisHelper.CleanRedisQuery()
		return res.Error
	} else {
		createData := topicModel.TopicClassify{
			ClassifyId: params.NewClassifyId,
			TopicId:    params.TopicId,
		}
		return midTopicClassifyCreate.CreateService(createData, db)
	}
}

