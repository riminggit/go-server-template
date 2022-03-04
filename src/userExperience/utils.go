package userExperience

import (
	"go-server-template/model/user"
	"go-server-template/pkg/level"
)

func JudgeUserCanAddExperience(data *userModel.UserExperience, experienceAdd int) *JudgeUserCanAddExperienceReturn {

	canAddExperience := &JudgeUserCanAddExperienceReturn{}
	levelMaxEx := level.PRIMARY_EXPERIENCE_UP

	if data.Level == level.MIDDLE_RANK {
		levelMaxEx = level.MIDDLE_RANK_EXPERIENCE_UP
	}
	if data.Level == level.ADVANCED {
		levelMaxEx = level.ADVANCED_EXPERIENCE_UP
	}

	if data.Level == level.SENIOR {
		levelMaxEx = level.SENIOR_EXPERIENCE_UP
	}

	if data.Level == level.SPECIALIST {
		levelMaxEx = level.SPECIALIST_EXPERIENCE_UP
	}

	if data.Level == level.SENIOR_SPECIALIST {
		levelMaxEx = level.SENIOR_SPECIALIST_EXPERIENCE_UP
	}

	if data.Level == level.RESEARCHER {
		levelMaxEx = level.RESEARCHER_EXPERIENCE_UP
	}

	canAddExperience.LevelMaxEx = levelMaxEx
	canAddExperience.CanAddEx = JudgeEXADD(levelMaxEx, data.Experience, experienceAdd)

	return canAddExperience
}

func Exalculate(params UpdateParams) int {
	experience := 0

	if params.Degree == 0 {
		experience = experience + level.EASY_EXPERIENCE
	} else {
		experience = experience + ((params.Degree + 1) * level.EASY_EXPERIENCE)
	}

	if params.IsBaseTopic == 0 {
		experience = experience + level.BASE_EXPERIENCE*2
	} else {
		experience = experience + level.BASE_EXPERIENCE
	}

	if params.IsImportantTopic != 0 {
		experience = experience + level.IMPORTANT_TOPIC
	}

	if params.Level != 0 {
		experience = experience + level.COEFFICIENT*params.Level
	}

	// 最大能获取经验为22
	return experience
}

func JudgeEXADD(levelExMax int, ex int, exAdd int) bool {
	res := true
	if levelExMax <= (ex + exAdd) {
		res = false
	}
	return res
}


func ExalculateTopicSet(params UpdateTopicSetParams) int {
	experience := 0

	if params.Level != 0 {
		experience = experience + level.TOPIC_SET_COEFFICIENT * params.Level
	}

	return experience
}