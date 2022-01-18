package userExperience

import (
	"github.com/gin-gonic/gin"
	"go-server-template/model/user"
	DB "go-server-template/pkg/db"
	"go-server-template/pkg/e"
	"go-server-template/pkg/utils"
)

func UserQueryExperience(c *gin.Context) *userQueryReturn {
	res := &userQueryReturn{}
	var queryInfo userModel.UserExperience
	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		res.Code = userInfoRes.Code
		return res
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", 1)
	queryFun = queryFun.Where("user_id = ?", userInfoRes.Data.ID)
	queryFun.Model(&userModel.UserExperience{}).Find(&queryInfo)

	res.Code = e.SUCCESS
	res.Data = queryInfo

	return res
}

// 用户查询自身经验信息，给其他接口调取用
func UserQueryExperienceSimple(c *gin.Context, userId int) *userQueryReturn {
	res := &userQueryReturn{}
	var queryInfo userModel.UserExperience

	queryFun := DB.DBLivingExample.Where("is_use = ?", 1)
	queryFun = queryFun.Where("user_id = ?", userId)
	err := queryFun.Model(&userModel.UserExperience{}).Find(&queryInfo).Error

	if err != nil {
		res.Code = e.NO_DATA_EXISTS
		return res
	}

	res.Code = e.SUCCESS
	res.Data = queryInfo

	return res
}

// 管理员查询用户信息
func QueryUserExperience(c *gin.Context, params QueryParams) *queryReturn {
	res := &queryReturn{}

	var queryInfo []userModel.UserExperience
	if params.IsUse == "" {
		params.IsUse = "1"
	}

	queryFun := DB.DBLivingExample.Where("is_use = ?", params.IsUse)
	if len(params.Experience) > 0 {
		queryFun = queryFun.Where("experience between ? and ?", params.Experience[0], params.Experience[1])
	}

	if params.Level != 0 {
		queryFun = queryFun.Where("level = ?", params.Level)
	}

	if params.UserId != 0 {
		queryFun = queryFun.Where("user_id = ?", params.UserId)
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

	queryFun = queryFun.Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize)
	queryFun.Model(&userModel.UserExperience{}).Find(&queryInfo).Count(&res.Data.PagingArgument.Total)

	res.Data.PagingArgument.PageNum = params.PageNum
	res.Data.PagingArgument.PageSize = params.PageSize
	res.Data.Data = queryInfo
	res.Code = e.SUCCESS

	return res
}
