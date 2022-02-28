package topicSetDelete

import (
	topicModel "go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	util "go-server-template/pkg/utils"

	"github.com/gin-gonic/gin"
)

func DeleteService(c *gin.Context, params DeleteParams) *DeleteReturn {
	res := &DeleteReturn{}

	if params.RealDel == "1" {
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Delete(&topicModel.TopicSet{}, "id IN ?", params.IDList).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.NameList) > 0 {
			delErr := DB.DBLivingExample.Delete(&topicModel.TopicSet{}, "name IN ?", params.NameList).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Data = append(res.Data, "名称条件删除失败")
			}
		}
	} else {
		setData := map[string]interface{}{"is_use": 0, "DeleteAt": util.GetNowTimeUnix()}
		if len(params.IDList) > 0 {
			delErr := DB.DBLivingExample.Model(&topicModel.TopicSet{}).Where("id IN ?", params.IDList).Updates(setData).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Data = append(res.Data, "id条件删除失败")
			}
		}
		if len(params.NameList) > 0 {
			delErr := DB.DBLivingExample.Model(&topicModel.TopicSet{}).Where("name IN ?", params.NameList).Updates(setData).Error
			if delErr != nil {
				logging.Error(delErr)
				res.Data = append(res.Data, "名称条件删除失败")
			}
		}
	}

	res.Code = e.SUCCESS
	return res
}
