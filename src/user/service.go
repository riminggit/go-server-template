package user

import (
	"encoding/json"
	"go-server-template/config"
	userModel "go-server-template/model/user"
	// "go-server-template/pkg/app"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	logging "go-server-template/pkg/log"
	Redis "go-server-template/pkg/redis"

	"github.com/gin-gonic/gin"
)

func QueryUserDataService(c *gin.Context, params QueryUserParams) *QueryUserReturnData {
	res := &QueryUserReturnData{}
	var userInfo []QueryUserData
	var RGetData UserReturnData

	dataRxpirationTime := projectConfig.AppConfig.BaseConfig.REDIS_COMMON_EXPIRATION_TIME

	redisParamsJson, _ := json.Marshal(params)
	interfaceName := "query-user-info:"
	queryRedisParams := interfaceName + string(redisParamsJson)

	redisData := Redis.GetValue(queryRedisParams)

	if redisData != "" {
		err := json.Unmarshal([]byte(redisData), &RGetData)
		if err != nil {
			logging.Debug(err)
		}
		res.Code = e.SUCCESS
		res.Data = RGetData
		return res
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)

	if params.Email != "" {
		queryFun = queryFun.Where("email = ?", params.Email)
	}

	if params.NickName != "" {
		queryFun = queryFun.Where("nick_name = ?", params.NickName)
	}

	if params.Phone != "" {
		queryFun = queryFun.Where("phone = ?", params.Phone)
	}

	if params.ComeFrom != "" {
		queryFun = queryFun.Where("come_from = ?", params.ComeFrom)
	}

	if params.Gender != "" {
		queryFun = queryFun.Where("gender = ?", params.Gender)
	}

	if len(params.City) > 0 {
		queryFun = queryFun.Where(map[string]interface{}{"city": params.City})
	}

	if len(params.Province) > 0 {
		queryFun = queryFun.Where(map[string]interface{}{"province": params.Province})
	}

	if len(params.Country) > 0 {
		queryFun = queryFun.Where(map[string]interface{}{"country": params.Country})
	}

	if len(params.Language) > 0 {
		queryFun = queryFun.Where(map[string]interface{}{"language": params.Language})
	}

	if len(params.Birthday) > 0 {
		queryFun = queryFun.Where("birthday between ? and ?", params.Birthday[0], params.Birthday[1])
	}

	if len(params.CreateAt) > 0 {
		queryFun = queryFun.Where("create_at between ? and ?", params.CreateAt[0], params.CreateAt[1])
	}

	if len(params.DeleteAt) > 0 {
		queryFun = queryFun.Where("delete_at between ? and ?", params.DeleteAt[0], params.DeleteAt[1])
	}

	if len(params.UpdateAt) > 0 {
		queryFun = queryFun.Where("update_at between ? and ?", params.UpdateAt[0], params.UpdateAt[1])
	}

	if params.IsAdmin != "" {
		queryFun = queryFun.Where("is_admin = ?", params.IsAdmin)
	}

	queryFun = queryFun.Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize)

	queryFun.Model(&userModel.User{}).Find(&userInfo).Count(&res.Data.PagingArgument.Total)

	res.Data.PagingArgument.PageNum = params.PageNum
	res.Data.PagingArgument.PageSize = params.PageSize
	res.Data.Data = userInfo

	RSetData := res.Data
	Redis.SetValue(queryRedisParams, RSetData, dataRxpirationTime)

	res.Code = e.SUCCESS

	return res
}
