package topicDelete

import (
	"github.com/gin-gonic/gin"
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/src/midTopicClassify/delete"
	"go-server-template/src/midTopicCompany/delete"
	"go-server-template/src/midTopicTag/delete"
	"go-server-template/src/midTopicType/delete"
)

func DeleteService(c *gin.Context, params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}
	res.Code = e.SUCCESS

	tx := DB.DBLivingExample.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			res.Code = e.CREATE_DATA_FILE
			res.Data = append(res.Data, "题目删除失败")
		}
	}()

	if len(params.IDList) > 0 {
		delErr := tx.Delete(&topicModel.Topic{}, "id IN ?", params.IDList).Error
		if delErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FILE
			return res
		}

		classifyParams := midTopicClassifyDelete.DeleteMultiple{
			TopicIdList: params.IDList,
		}
		delClassifyErr := midTopicClassifyDelete.DeleteMultipleService(classifyParams, tx)
		if delClassifyErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FILE
		}

		companyParams := midTopicCompanyDelete.DeleteMultiple{
			TopicIdList: params.IDList,
		}
		delCompanyErr := midTopicCompanyDelete.DeleteMultipleService(companyParams, tx)
		if delCompanyErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FILE
		}

		tagParams := midTopicTagDelete.DeleteMultiple{
			TopicIdList: params.IDList,
		}
		delTagErr := midTopicTagDelete.DeleteMultipleService(tagParams, tx)
		if delTagErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FILE
		}

		typeParams := midTopicTypeDelete.DeleteMultiple{
			TopicIdList: params.IDList,
		}
		delTypeErr := midTopicTypeDelete.DeleteMultipleService(typeParams, tx)
		if delTypeErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FILE
		}

		commitErr := tx.Commit().Error
		if commitErr != nil {
			res.Code = e.CREATE_DATA_FILE
		}
		return res
	}

	return res
}
