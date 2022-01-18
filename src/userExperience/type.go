package userExperience

import (
	"go-server-template/model/user"
)

type UpdateParams struct {
	Degree           int `json:"degree"`             // 难度，简单：0，中等：1，难：2，极难：3
	Level            int `json:"level"`              // 1 初级 ， 2 中级 ， 3 高级 ， 4 资深 ，5 专家 ， 6 资深专家 ， 7 研究员
	IsBaseTopic      int `json:"is_base_topic" `     // 0 否  1是 ，是否是基础题
	IsImportantTopic int `json:"is_important_topic"` // 0 否  1是 ，是否是重点题
}

type UpdateReturn struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type QueryParams struct {
	UserId     int      `json:"user_id" `   // 用户id
	Experience []int    `json:"experience"` // 用户经验
	Level      int      `json:"level"`      // 用户等级
	CreateAt   []string `json:"create_at"`
	DeleteAt   []string `json:"delete_at"`
	UpdateAt   []string `json:"update_at"`
	IsUse      string   `json:"is_use"`
	PageNum    int      `json:"pageNum"`
	PageSize   int      `json:"pageSize"`
}

type queryReturn struct {
	Code int             `json:"code"`
	Data queryReturnData `json:"data"`
}

type queryReturnData struct {
	PagingArgument PageArgument               `json:"pageArgument"`
	Data           []userModel.UserExperience `json:"data"`
}

type userQueryReturn struct {
	Code int                      `json:"code"`
	Data userModel.UserExperience `json:"data"`
}

type JudgeUserCanAddExperienceReturn struct {
	LevelMaxEx int  `json:"level_max_ex"`
	CanAddEx   bool `json:"can_add_ex"`
}
type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}
