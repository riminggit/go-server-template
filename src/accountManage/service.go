package account

import (
	"encoding/json"
	userModel "go-server-template/model/user"
	"go-server-template/pkg/app"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"

	"github.com/gin-gonic/gin"
)

func ChangePhoneService(c *gin.Context, params ChangePhoneParams) *ChangePhoneReturnData {

	res := &ChangePhoneReturnData{}
	token := app.GetHeaderToken(c)

	var userInfoRedis userModel.User
	if token != "" {
		getUserInfoerr := json.Unmarshal([]byte(Redis.GetValue(token)), &userInfoRedis)
		if getUserInfoerr != nil {
			logging.Debug(getUserInfoerr)
		}
	} else if token == "" {
		res.Code = e.NO_DATA_EXISTS
		return res
	}

	DB.DBLivingExample.Table("user").Where("id = ?", userInfoRedis.ID).Updates(params)

	res.Code = e.SUCCESS

	return res
}
