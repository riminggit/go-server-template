package midTopicTagUpdate

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicTag/create"
	"go-server-template/src/midTopicTag/helper"
	"time"
	"gorm.io/gorm"
)

func UpdateService(params UpdateParams, db *gorm.DB) error {

	var queryInfo []topicModel.TopicTag

	queryFun := db
	queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	queryFun = queryFun.Where("tag_id = ?", params.TagId)
	queryFun.Model(&topicModel.TopicTag{}).Find(&queryInfo)

	if len(queryInfo) > 0 {
		setData := topicModel.TopicTag{
			UpdateAt: time.Now().Add(8 * time.Hour),
			TagId:    params.NewTagId,
		}
		res := queryFun.Updates(setData)
		thisHelper.CleanRedisQuery()
		return res.Error
	} else {
		createData := topicModel.TopicTag{
			TagId:  params.NewTagId,
			TopicId: params.TopicId,
		}
		return midTopicTagCreate.CreateService(createData, db)
	}
}
