package util

import (
	"encoding/json"
	// "fmt"
	"github.com/gin-gonic/gin"
	userModel "go-server-template/model/user"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
)

type userInfoReturn struct {
	Code int
	Data userModel.User
}

func GetUserInfo(c *gin.Context) *userInfoReturn {
	res := &userInfoReturn{}
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

	res.Data = userInfoRedis
	res.Code = e.SUCCESS
	return res
}
