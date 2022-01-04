package midTopicClassifyDelete

import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/midTopicClassify/helper"
)

func DeleteService(params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}

	if params.ID != "" {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicClassify{}, "id = ?", params.ID).Error
		if delErr != nil {
			res.Data = append(res.Data, "id条件删除失败")
		}
	}

	if params.TopicId != "" {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicClassify{}, "topic_id = ?", params.TopicId).Error
		if delErr != nil {
			res.Data = append(res.Data, "题目id条件删除失败")
		}
	}

	if params.ClassifyId != "" {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicClassify{}, "classify_id = ?", params.ClassifyId).Error
		if delErr != nil {
			res.Data = append(res.Data, "分类id条件删除失败")
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
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicClassify{}, "id IN ?", params.IDList).Error
		if delErr != nil {
			res.Data = append(res.Data, "id条件删除失败")
		}
	}
	if len(params.TopicIdList) > 0 {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicClassify{}, "topic_id IN ?", params.TopicIdList).Error
		if delErr != nil {
			res.Data = append(res.Data, "题目id条件删除失败")
		}
	}
	if len(params.ClassifyIdList) > 0 {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicClassify{}, "classify_id IN ?", params.ClassifyIdList).Error
		if delErr != nil {
			res.Data = append(res.Data, "分类id条件删除失败")
		}
	}

	// 清除查询的redis缓存
	thisHelper.CleanRedisQuery()

	res.Code = e.SUCCESS
	return res
}
