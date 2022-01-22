package sections

import (
	"go-server-template/pkg/e"
	"go-server-template/pkg/level"
	util "go-server-template/pkg/utils"
	"go-server-template/src/userExperience"
	"strconv"

	"github.com/gin-gonic/gin"
)


func JudgeQueryCondition(c *gin.Context, queryCount string, table string, queryWhere string) string {
	userInfoRes := util.GetUserInfo(c)
	if userInfoRes.Code != e.SUCCESS {
		return ""
	}
	userExperience := userExperience.UserQueryExperienceSimple(c, userInfoRes.Data.ID)
	if userExperience.Code != e.SUCCESS {
		return ""
	}

	var condition string
	num := util.RandNum()
	if userExperience.Data.Level == level.PRIMARY {
		if num >= 0 && num <= 50 {
			condition = strconv.Itoa(level.PRIMARY) // 50
		} else if num <= 75 {
			condition = strconv.Itoa(level.MIDDLE_RANK) // 25
		} else if num <= 90 {
			condition = strconv.Itoa(level.ADVANCED) //15
		} else if num <= 95 {
			condition = strconv.Itoa(level.SENIOR) // 5
		} else if num <= 98 {
			condition = strconv.Itoa(level.SPECIALIST) // 3
		} else {
			condition = strconv.Itoa(level.SENIOR_SPECIALIST) // 1
		}
	}
	if userExperience.Data.Level == level.MIDDLE_RANK {
		if num >= 0 && num <= 40 {
			condition = strconv.Itoa(level.PRIMARY) // 40
		} else if num <= 70 {
			condition = strconv.Itoa(level.MIDDLE_RANK) // 30
		} else if num <= 90 {
			condition = strconv.Itoa(level.ADVANCED) // 20
		} else if num <= 95 {
			condition = strconv.Itoa(level.SENIOR) // 5
		} else if num <= 98 {
			condition = strconv.Itoa(level.SPECIALIST) // 3
		} else {
			condition = strconv.Itoa(level.SENIOR_SPECIALIST) // 1
		}
	}
	if userExperience.Data.Level == level.ADVANCED {
		if num >= 0 && num <= 25 {
			condition = strconv.Itoa(level.PRIMARY) // 25
		} else if num <= 50 {
			condition = strconv.Itoa(level.MIDDLE_RANK) // 25
		} else if num <= 85 {
			condition = strconv.Itoa(level.ADVANCED) // 35
		} else if num <= 93 {
			condition = strconv.Itoa(level.SENIOR) // 8
		} else if num <= 98 {
			condition = strconv.Itoa(level.SPECIALIST) // 5
		} else {
			condition = strconv.Itoa(level.SENIOR_SPECIALIST) // 1
		}
	}

	if userExperience.Data.Level == level.SENIOR {
		if num >= 0 && num <= 20 {
			condition = strconv.Itoa(level.PRIMARY) // 20
		} else if num <= 45 {
			condition = strconv.Itoa(level.MIDDLE_RANK) // 25
		} else if num <= 70 {
			condition = strconv.Itoa(level.ADVANCED) // 28
		} else if num <= 88 {
			condition = strconv.Itoa(level.SENIOR) // 18
		} else if num <= 95 {
			condition = strconv.Itoa(level.SPECIALIST) // 8
		} else {
			condition = strconv.Itoa(level.SENIOR_SPECIALIST) // 4
		}
	}

	if userExperience.Data.Level == level.SPECIALIST {
		if num >= 0 && num <= 10 {
			condition = strconv.Itoa(level.PRIMARY) // 10
		} else if num <= 30 {
			condition = strconv.Itoa(level.MIDDLE_RANK) // 20
		} else if num <= 55 {
			condition = strconv.Itoa(level.ADVANCED) // 25
		} else if num <= 85 {
			condition = strconv.Itoa(level.SENIOR) // 30
		} else if num <= 95 {
			condition = strconv.Itoa(level.SPECIALIST) // 10
		} else if num <= 98 {
			condition = strconv.Itoa(level.SENIOR_SPECIALIST) // 3
		} else {
			condition = strconv.Itoa(level.RESEARCHER) // 1
		}
	}

	if userExperience.Data.Level == level.SENIOR_SPECIALIST {
		if num >= 0 && num <= 5 {
			condition = strconv.Itoa(level.PRIMARY) // 5
		} else if num <= 20 {
			condition = strconv.Itoa(level.MIDDLE_RANK) // 15
		} else if num <= 40 {
			condition = strconv.Itoa(level.ADVANCED) // 20
		} else if num <= 65 {
			condition = strconv.Itoa(level.SENIOR) // 25
		} else if num <= 90 {
			condition = strconv.Itoa(level.SPECIALIST) // 25
		} else if num <= 96 {
			condition = strconv.Itoa(level.SENIOR_SPECIALIST) // 6
		} else {
			condition = strconv.Itoa(level.RESEARCHER) // 4
		}
	}

	if userExperience.Data.Level == level.RESEARCHER {
		if num >= 0 && num <= 3 {
			condition = strconv.Itoa(level.PRIMARY) // 3
		} else if num <= 15 {
			condition = strconv.Itoa(level.MIDDLE_RANK) // 12
		} else if num <= 30 {
			condition = strconv.Itoa(level.ADVANCED) // 15
		} else if num <= 50 {
			condition = strconv.Itoa(level.SENIOR) // 20
		} else if num <= 85 {
			condition = strconv.Itoa(level.SPECIALIST) // 35
		} else if num <= 95 {
			condition = strconv.Itoa(level.SENIOR_SPECIALIST) // 10
		} else {
			condition = strconv.Itoa(level.RESEARCHER) // 5
		}
	}

	if condition != "" {
		// 根据条件随机查询
		return "SELECT * FROM " + table + " WHERE " + queryWhere + " = " + condition + " ORDER BY RAND() LIMIT " + queryCount
	} else {
		return ""
	}

}
