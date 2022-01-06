package midTopicTagDelete

import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/midTopicTag/helper"
)

func DeleteService(params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}

	if params.ID != "" {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicTag{}, "id = ?", params.ID).Error
		if delErr != nil {
			res.Data = append(res.Data, "id条件删除失败")
		}
	}

	if params.TopicId != "" {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicTag{}, "topic_id = ?", params.TopicId).Error
		if delErr != nil {
			res.Data = append(res.Data, "题目id条件删除失败")
		}
	}

	if params.TagId != "" {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicTag{}, "tag_id = ?", params.TagId).Error
		if delErr != nil {
			res.Data = append(res.Data, "标签id条件删除失败")
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
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicTag{}, "id IN ?", params.IDList).Error
		if delErr != nil {
			res.Data = append(res.Data, "id条件删除失败")
		}
	}
	if len(params.TopicIdList) > 0 {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicTag{}, "topic_id IN ?", params.TopicIdList).Error
		if delErr != nil {
			res.Data = append(res.Data, "题目id条件删除失败")
		}
	}
	if len(params.TagIdList) > 0 {
		delErr := DB.DBLivingExample.Delete(&topicModel.TopicTag{}, "tag_id IN ?", params.TagIdList).Error
		if delErr != nil {
			res.Data = append(res.Data, "标签id条件删除失败")
		}
	}

	// 清除查询的redis缓存
	thisHelper.CleanRedisQuery()

	res.Code = e.SUCCESS
	return res
}
