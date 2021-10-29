package userLogin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-server-template/model/user"
	"go-server-template/pkg/app"
	"go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/pkg/log"
	"go-server-template/pkg/redis"
	"go-server-template/pkg/util"
	// "go-server-template/src/wxDispose"   // 直接拿参数解密
	"strconv"
	"time"
)

func UserWXLoginService(c *gin.Context, params WXUserParams) *LoginReturnData {
	res := &LoginReturnData{}
	token := app.GetHeaderToken(c)

	if token != "" {
		userInfo := Redis.GetValue(token)
		if userInfo != "" {
			userInfoJsonData := &userModel.User{}
			err := json.Unmarshal([]byte(userInfo), userInfoJsonData)
			if err != nil {
				logging.Debug(err)
			}
			res.Code = e.IS_LOGIN
			res.Token = token
			return res
		}
	}

	// 判断结构体如果为空
	if params == (WXUserParams{}) {
		res.Code = e.INVALID_PARAMS
		return res
	} else {

		// 将一个结构体的值给另一个
		createParams := WXUserCreateParams{}
		util.StructAssign(&params, &createParams)

		// getDecodeDataParam := wxDispose.WXDncryptParams{}
		// util.StructAssign(&params, &getDecodeDataParam)

		// decodeData, decodeErr := wxDispose.WXDncryptService(getDecodeDataParam)
		// if decodeErr != nil {
		// 	res.Code = e.WX_DNCRYPT_FAIL
		// }

		// 查询用户信息
		var userInfo userModel.User
		// 获取第一条匹配的记录
		err := DB.DBLivingExample.Table("user").Where("openid = ?", params.Openid).First(&userInfo).Error
		if err != nil {
			result := DB.DBLivingExample.Table("user").Create(createParams)
			if result.Error != nil {
				logging.Error(result.Error)
				res.Code = e.CREATE_DATA_FILE
				return res
			}
		} else {
			DB.DBLivingExample.Table("user").Where("openid = ?", params.Openid).Updates(createParams)
		}

		// 生成token
		newToken, tokenErr := util.GenerateToken(userInfo.NickName, userInfo.Openid)
		if tokenErr != nil {
			res.Code = e.ERROR_AUTH_TOKEN
			logging.Error(tokenErr)
			return res
		}

		// 将用户信息存入redis
		Redis.SetValue(newToken, userInfo, 60)

		if strconv.Itoa(userInfo.ID) == "" {
			queryErr := DB.DBLivingExample.Table("user").Where("openid = ?", params.Openid).First(&userInfo).Error
			if err != nil {
				logging.Error(queryErr)
			}
		}

		// 更新用户登陆信息记录表
		var recordParam userModel.UserLoginRecord
		recordParam.ComeFrom = "WX"
		recordParam.LoginAt = time.Now()
		recordParam.UserId = userInfo.ID

		recordCreateResult := DB.DBLivingExample.Table("user_login_record").Create(&recordParam)
		if recordCreateResult.Error != nil {
			logging.Error(recordCreateResult.Error)
		}

		res.Token = newToken
		res.Code = e.SUCCESS

		return res
	}

}
