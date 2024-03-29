package userLoginAndLayout

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	userModel "go-server-template/model/user"
	"go-server-template/pkg/app"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"
	"go-server-template/pkg/util"
	"strconv"
	"time"
)

func UserWXLoginService(c *gin.Context, params WXUserCreateParams) *LoginReturnData {
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
	if params == (WXUserCreateParams{}) {
		res.Code = e.INVALID_PARAMS
		return res
	} else {

		// 查询用户信息
		var userInfo userModel.User
		// 获取第一条匹配的记录
		err := DB.DBLivingExample.Table("user").Where("openid = ?", params.Openid).First(&userInfo).Error
		if err != nil {
			result := DB.DBLivingExample.Table("user").Create(params)
			if result.Error != nil {
				logging.Error(result.Error)
				res.Code = e.CREATE_DATA_FILE
				return res
			}
		}
		// } else {
		// 更新同步
		// DB.DBLivingExample.Table("user").Where("openid = ?", params.Openid).Updates(params)
		// }

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
			if queryErr != nil {
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

func LoginService(c *gin.Context, params LoginParams) *LoginReturnData {
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
	if params == (LoginParams{}) {
		res.Code = e.INVALID_PARAMS
		return res
	} else {
		// 查询用户信息
		var userInfo userModel.User
		switch params.LoginType {
		case "email":
			err := DB.DBLivingExample.Table("user").Where("email = ?", params.Email).First(&userInfo).Error
			if err != nil {
				logging.Error(err)
				res.Code = e.EMAIL_DOES_NOT_EXIST
				return res
			}

		case "phone":
			err := DB.DBLivingExample.Table("user").Where("phone = ?", params.Phone).First(&userInfo).Error
			if err != nil {
				logging.Error(err)
				res.Code = e.PHONE_DOES_NOT_EXIST
				return res
			}

		case "nick_name":
			err := DB.DBLivingExample.Table("user").Where("nick_name = ?", params.NickName).First(&userInfo).Error
			if err != nil {
				logging.Error(err)
				res.Code = e.NICKNAME_DOES_NOT_EXIST
				return res
			}

		default:
			res.Code = e.ACCOUNT_DOES_NOT_EXIST
			return res
		}

		if userInfo.Password != params.Password {
			res.Code = e.PASSWORD_MISTAKE
			return res
		}

		// 生成token
		newToken, tokenErr := util.GenerateToken(userInfo.NickName, userInfo.Password)
		if tokenErr != nil {
			res.Code = e.ERROR_AUTH_TOKEN
			logging.Error(tokenErr)
			return res
		}

		// 将用户信息存入redis
		Redis.SetValue(newToken, userInfo, 60)

		recordParam := &userModel.UserLoginRecord{
			ComeFrom: params.ComeFrom,
			LoginAt:  time.Now(),
			UserId:   userInfo.ID,
		}

		recordCreateResult := DB.DBLivingExample.Table("user_login_record").Create(&recordParam)
		if recordCreateResult.Error != nil {
			logging.Error(recordCreateResult.Error)
		}

		res.Token = newToken
		res.Code = e.SUCCESS
		return res
	}

}
