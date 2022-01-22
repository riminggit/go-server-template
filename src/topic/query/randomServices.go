package topicQuery

import (
	"fmt"
	"go-server-template/model/topic"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	util "go-server-template/pkg/utils"
	"go-server-template/src/sections"

	"github.com/gin-gonic/gin"
)

func QueryTopicRandomService(c *gin.Context, params QueryTopicRandomParams) *QueryTopicRandomReturn {
	res := &QueryTopicRandomReturn{}
	res.Code = e.SUCCESS
	var queryInfo topicModel.Topic
	var querySql string
	queryCount := "1"
	tableName := `topic`

	querySqlHelper := util.RandSql(tableName, queryCount)

	if params.IsQueryByUserLevel == "1" {
		sqlHelper := sections.JudgeQueryCondition(c, queryCount, tableName, "level")
		if sqlHelper != "" {
			querySql = sqlHelper
		} else {
			querySql = querySqlHelper
		}
	} else {
		querySql = querySqlHelper
	}

	fmt.Println(querySql, "querySql")
	queryFun := DB.DBLivingExample.Raw(querySql).Scan(&queryInfo)

	// a := DB.DBLivingExample.Model(&topicModel.Topic{}).

	if queryFun.Error != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}
	if queryInfo.ID == 0 {
		querySql = querySqlHelper
		queryFunAgain := DB.DBLivingExample.Raw(querySql).Scan(&queryInfo)
		if queryFunAgain.Error != nil {
			res.Code = e.NO_DATA_EXISTS
			return res
		}
	}

	fmt.Println(queryInfo, "querySqlHelper")

	res.Data = queryInfo

	return res
}
