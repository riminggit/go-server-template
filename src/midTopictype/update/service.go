package midTopicTypeUpdate


import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/midTopicType/create"
	"go-server-template/src/midTopicType/helper"
	"strconv"
	"time"
)

func UpdateService(params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}

	var resData []string
	var queryInfo []topicModel.TopicType

	queryFun := DB.DBLivingExample
	queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	queryFun = queryFun.Where("type_id = ?", params.TypeId)
	queryFun.Model(&topicModel.TopicType{}).Find(&queryInfo)

	if len(queryInfo) > 0 {
		setData := topicModel.TopicType{
			UpdateAt:   time.Now().Add(8 * time.Hour),
			TypeId: params.NewTypeId,
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
		createData := midTopicTypeCreate.CreateParams{
			TypeId: params.NewTypeId,
			TopicId:    params.TopicId,
		}
		res := midTopicTypeCreate.CreateService(createData)
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
	var queryInfo []topicModel.TopicType
	for index, item := range params.Data {

		queryFun := DB.DBLivingExample
		queryFun = queryFun.Where("topic_id = ?", item.TopicId)
		queryFun = queryFun.Where("type_id = ?", item.TypeId)
		queryFun.Model(&topicModel.TopicType{}).Find(&queryInfo)

		if len(queryInfo) > 0 {
			setData := topicModel.TopicType{
				UpdateAt:   time.Now().Add(8 * time.Hour),
				TypeId: item.NewTypeId,
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
			createData := midTopicTypeCreate.CreateParams{
				TypeId: item.NewTypeId,
				TopicId:    item.TopicId,
			}
			res := midTopicTypeCreate.CreateService(createData)
			resData = append(resData, res.Data...)
		}

	}

	// 清除查询的redis缓存
	thisHelper.CleanRedisQuery()

	res.Data = resData
	res.Code = e.SUCCESS
	return res
}
