package midTopicCompanyDelete

import (
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/midTopicCompany/create"
	"go-server-template/src/midTopicCompany/helper"
	"strconv"
	"time"
)

func UpdateService(params UpdateParams) *UpdateReturn {
	res := &UpdateReturn{}

	var resData []string
	var queryInfo []topicModel.TopicCompany

	queryFun := DB.DBLivingExample
	queryFun = queryFun.Where("topic_id = ?", params.TopicId)
	queryFun = queryFun.Where("company_id = ?", params.CompanyId)
	queryFun.Model(&topicModel.TopicCompany{}).Find(&queryInfo)

	if len(queryInfo) > 0 {
		setData := topicModel.TopicCompany{
			UpdateAt:   time.Now().Add(8 * time.Hour),
			CompanyId: params.NewCompanyId,
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
		createData := midTopicCompanyCreate.CreateParams{
			CompanyId: params.NewCompanyId,
			TopicId:    params.TopicId,
		}
		res := midTopicCompanyCreate.CreateService(createData)
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
	var queryInfo []topicModel.TopicCompany
	for index, item := range params.Data {

		queryFun := DB.DBLivingExample
		queryFun = queryFun.Where("topic_id = ?", item.TopicId)
		queryFun = queryFun.Where("company_id = ?", item.CompanyId)
		queryFun.Model(&topicModel.TopicCompany{}).Find(&queryInfo)

		if len(queryInfo) > 0 {
			setData := topicModel.TopicCompany{
				UpdateAt:   time.Now().Add(8 * time.Hour),
				CompanyId: item.NewCompanyId,
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
			createData := midTopicCompanyCreate.CreateParams{
				CompanyId: item.NewCompanyId,
				TopicId:    item.TopicId,
			}
			res := midTopicCompanyCreate.CreateService(createData)
			resData = append(resData, res.Data...)
		}

	}

	// 清除查询的redis缓存
	thisHelper.CleanRedisQuery()

	res.Data = resData
	res.Code = e.SUCCESS
	return res
}
