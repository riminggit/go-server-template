package midTopicClassifyUpdate

import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/midTopicClassify/create"
	"go-server-template/src/midTopicClassify/helper"
	"strconv"
	"time"
)

func UpdateService(params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}

	var resData []string
	var queryInfo []topicModel.TopicClassify

	queryFun := DB.DBLivingExample
	queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	queryFun = queryFun.Where("classify_id = ?", params.ClassifyId)
	queryFun.Model(&topicModel.TopicClassify{}).Find(&queryInfo)

	if len(queryInfo) > 0 {
		setData := topicModel.TopicClassify{
			UpdateAt:   time.Now().Add(8 * time.Hour),
			ClassifyId: params.NewClassifyId,
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
		createData := midTopicClassifyCreate.CreateParams{
			ClassifyId: params.NewClassifyId,
			TopicId:    params.TopicId,
		}
		res := midTopicClassifyCreate.CreateService(createData)
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
	var queryInfo []topicModel.TopicClassify
	for index, item := range params.Data {

		queryFun := DB.DBLivingExample
		queryFun = queryFun.Where("topic_id = ?", item.TopicId)
		queryFun = queryFun.Where("classify_id = ?", item.ClassifyId)
		queryFun.Model(&topicModel.TopicClassify{}).Find(&queryInfo)

		if len(queryInfo) > 0 {
			setData := topicModel.TopicClassify{
				UpdateAt:   time.Now().Add(8 * time.Hour),
				ClassifyId: item.NewClassifyId,
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
			createData := midTopicClassifyCreate.CreateParams{
				ClassifyId: item.NewClassifyId,
				TopicId:    item.TopicId,
			}
			res := midTopicClassifyCreate.CreateService(createData)
			resData = append(resData, res.Data...)
		}

	}

	// 清除查询的redis缓存
	thisHelper.CleanRedisQuery()

	res.Data = resData
	res.Code = e.SUCCESS
	return res
}
