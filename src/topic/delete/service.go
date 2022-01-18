package topicDelete

import (
	topicModel "go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	midTopicClassifyDelete "go-server-template/src/midTopicClassify/delete"
	midTopicCompanyDelete "go-server-template/src/midTopicCompany/delete"
	midTopicTagDelete "go-server-template/src/midTopicTag/delete"
	midTopicTypeDelete "go-server-template/src/midTopicType/delete"
	"github.com/gin-gonic/gin"
)

func DeleteService(c *gin.Context, params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}
	res.Code = e.SUCCESS

	tx := DB.DBLivingExample.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			res.Code = e.CREATE_DATA_FALE
			res.Data = append(res.Data, "题目删除失败")
		}
	}()

	if len(params.IDList) > 0 {
		delErr := tx.Delete(&topicModel.Topic{}, "id IN ?", params.IDList).Error
		if delErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FALE
			return res
		}

		classifyParams := midTopicClassifyDelete.DeleteMultiple{
			TopicIdList: params.IDList,
		}
		delClassifyErr := midTopicClassifyDelete.DeleteMultipleService(classifyParams, tx)
		if delClassifyErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FALE
		}

		companyParams := midTopicCompanyDelete.DeleteMultiple{
			TopicIdList: params.IDList,
		}
		delCompanyErr := midTopicCompanyDelete.DeleteMultipleService(companyParams, tx)
		if delCompanyErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FALE
		}

		tagParams := midTopicTagDelete.DeleteMultiple{
			TopicIdList: params.IDList,
		}
		delTagErr := midTopicTagDelete.DeleteMultipleService(tagParams, tx)
		if delTagErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FALE
		}

		typeParams := midTopicTypeDelete.DeleteMultiple{
			TopicIdList: params.IDList,
		}
		delTypeErr := midTopicTypeDelete.DeleteMultipleService(typeParams, tx)
		if delTypeErr != nil {
			tx.Rollback()
			res.Code = e.DELETE_DATA_FALE
		}

		commitErr := tx.Commit().Error
		if commitErr != nil {
			res.Code = e.CREATE_DATA_FALE
		}
		return res
	}

	return res
}
