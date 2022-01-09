package util

import (
	"encoding/json"
	"strconv"
	"go-server-template/model/user"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	"github.com/gin-gonic/gin"
)

type UserInfoReturn struct {
	Code int
	Data userModel.User
}

type UserPrefixReturn struct {
	UserData userModel.User
	UserPrefix string
}

func GetUserInfo(c *gin.Context) *UserInfoReturn {
	res := &UserInfoReturn{}
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


func GetUserPrefix(c *gin.Context) *UserPrefixReturn {
	res:= &UserPrefixReturn{}
	userInfoRes := GetUserInfo(c)
	userId := userInfoRes.Data.ID
	// int转string
	res.UserPrefix =  "user:" + strconv.Itoa(userId)
	res.UserData = userInfoRes.Data
	return res
}
