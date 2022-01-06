package midTopicTagUpdate

import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/midTopicTag/create"
	"go-server-template/src/midTopicTag/helper"
	"strconv"
	"time"
)

func UpdateService(params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}

	var resData []string
	var queryInfo []topicModel.TopicTag

	queryFun := DB.DBLivingExample
	queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	queryFun = queryFun.Where("tag_id = ?", params.TagId)
	queryFun.Model(&topicModel.TopicTag{}).Find(&queryInfo)

	if len(queryInfo) > 0 {
		setData := topicModel.TopicTag{
			UpdateAt:   time.Now().Add(8 * time.Hour),
			TagId: params.NewTagId,
		}
		res := queryFun.Updates(setData)
		if res.Error != nil {
			msg := "数据更新出错"
			resData = append(resData, msg)
		} else {
			msg := "数据更新成功"
			resData = append(resData, msg)
		}
	} else {
		createData := midTopicTagCreate.CreateParams{
			TagId: params.NewTagId,
			TopicId:    params.TopicId,
		}
		res := midTopicTagCreate.CreateService(createData)
		resData = append(resData, res.Data...)
	}

	// 清除查询的redis缓存
	thisHelper.CleanRedisQuery()

	res.Data = resData
	res.Code = e.SUCCESS
	return res
}

func UpdateMultipleService(params UpdateMultiple) *UpdateReturn {
	res := &UpdateReturn{}

	var resData []string
	var queryInfo []topicModel.TopicTag
	for index, item := range params.Data {

		queryFun := DB.DBLivingExample
		queryFun = queryFun.Where("topic_id = ?", item.TopicId)
		queryFun = queryFun.Where("tag_id = ?", item.TagId)
		queryFun.Model(&topicModel.TopicTag{}).Find(&queryInfo)

		if len(queryInfo) > 0 {
			setData := topicModel.TopicTag{
				UpdateAt:   time.Now().Add(8 * time.Hour),
				TagId: item.NewTagId,
			}
			res := queryFun.Updates(setData)
			if res.Error != nil {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据更新出错"
				resData = append(resData, msg)
			} else {
				msg := "第" + strconv.Itoa((index + 1)) + "条数据更新成功"
				resData = append(resData, msg)
			}
		} else {
			createData := midTopicTagCreate.CreateParams{
				TagId: item.NewTagId,
				TopicId:    item.TopicId,
			}
			res := midTopicTagCreate.CreateService(createData)
			resData = append(resData, res.Data...)
		}

	}

	// 清除查询的redis缓存
	thisHelper.CleanRedisQuery()

	res.Data = resData
	res.Code = e.SUCCESS
	return res
}
