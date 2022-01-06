package midTopicTypeDelete

import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/midTopicType/helper"
)

func DeleteService(params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}

	if params.ID != "" {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicType{}, "id = ?", params.ID).Error
		if delErr != nil {
			res.Data = append(res.Data, "id条件删除失败")
		}
	}

	if params.TopicId != "" {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicType{}, "topic_id = ?", params.TopicId).Error
		if delErr != nil {
			res.Data = append(res.Data, "题目id条件删除失败")
		}
	}

	if params.TypeId != "" {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicType{}, "type_id = ?", params.TypeId).Error
		if delErr != nil {
			res.Data = append(res.Data, "二级分类id条件删除失败")
		}
	}

	// 清除查询的redis缓存
	thisHelper.CleanRedisQuery()

	res.Code = e.SUCCESS
	return res
}

func DeleteMultipleService(params DeleteMultiple) *DeleteReturn {
	res := &DeleteReturn{}

	if len(params.IDList) > 0 {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicType{}, "id IN ?", params.IDList).Error
		if delErr != nil {
			res.Data = append(res.Data, "id条件删除失败")
		}
	}
	if len(params.TopicIdList) > 0 {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicType{}, "topic_id IN ?", params.TopicIdList).Error
		if delErr != nil {
			res.Data = append(res.Data, "题目id条件删除失败")
		}
	}
	if len(params.TypeIdList) > 0 {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicType{}, "type_id IN ?", params.TypeIdList).Error
		if delErr != nil {
			res.Data = append(res.Data, "二级分类id条件删除失败")
		}
	}

	// 清除查询的redis缓存
	thisHelper.CleanRedisQuery()

	res.Code = e.SUCCESS
	return res
}
