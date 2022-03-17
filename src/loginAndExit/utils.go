package userloginAndExit

import (
	"encoding/json"
	userModel "go-server-template/model/user"
	"go-server-template/pkg/app"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"

	"github.com/gin-gonic/gin"
)

func CheckLogin(c *gin.Context) *LoginReturnData {
	res := &LoginReturnData{
		Code: e.USER_NOT_LOGIN,
	}
	token := app.GetHeaderToken(c)

	if token != "" {
		userInfoRedis := Redis.GetValue(token)
		if userInfoRedis != "" {
			userInfoJsonData := &userModel.User{}
			err := json.Unmarshal([]byte(userInfoRedis), userInfoJsonData)
			if err != nil {
				logging.Debug(err)
			}
			res.Code = e.IS_LOGIN
			res.Token = token
			res.UserInfo = *userInfoJsonData
		}
	}
	return res
}

func CheckNeedUpdate(c *gin.Context, userInfo *userModel.User, params *userModel.User) bool {
	res := false

	if params.AvatarUrl != "" && params.AvatarUrl != userInfo.AvatarUrl {
		res = true
	}
	if params.ComeFrom != "" && params.ComeFrom != userInfo.ComeFrom {
		res = true
	}
	if params.NickName != "" && params.NickName != userInfo.NickName {
		res = true
	}
	if params.Gender != userInfo.Gender {
		res = true
	}
	if params.City != "" && params.City != userInfo.City {
		res = true
	}
	if params.Province != "" && params.Province != userInfo.Province {
		res = true
	}
	if params.Country != "" && params.Country != userInfo.Country {
		res = true
	}
	if params.Language != "" && params.Language != userInfo.Language {
		res = true
	}
	if params.Rawdata != "" && params.Rawdata != userInfo.Rawdata {
		res = true
	}
	if params.Signature != "" && params.Signature != userInfo.Signature {
		res = true
	}
	if params.Iv != "" && params.Iv != userInfo.Iv {
		res = true
	}

	return res
}
