package midTopicTypeUpdate

import (
	"go-server-template/model/topic"
	"go-server-template/src/midTopicType/create"
	"go-server-template/src/midTopicType/helper"
	"gorm.io/gorm"
	"time"
)

func UpdateService(params UpdateParams, db *gorm.DB) error {

	var queryInfo []topicModel.TopicType
	queryFun := db
	queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	queryFun = queryFun.Where("type_id = ?", params.TypeId)
	queryFun.Model(&topicModel.TopicType{}).Find(&queryInfo)

	if len(queryInfo) > 0 {
		setData := topicModel.TopicType{
			UpdateAt: time.Now().Add(8 * time.Hour),
			TypeId:   params.NewTypeId,
		}
		res := queryFun.Updates(setData)
		thisHelper.CleanRedisQuery()
		return res.Error
	} else {
		createData := midTopicTypeCreate.CreateParams{
			TypeId:  params.NewTypeId,
			TopicId: params.TopicId,
		}
		return midTopicTypeCreate.CreateService(createData, db)
	}

}

