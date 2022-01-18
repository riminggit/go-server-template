package level

var LevelName = map[int]string{
	PRIMARY:           "初级",
	MIDDLE_RANK:       "中级",
	ADVANCED:          "高级",
	SENIOR:            "资深",
	SPECIALIST:        "专家",
	SENIOR_SPECIALIST: "资深专家",
	RESEARCHER:        "研究员",
}

func GetLevelName(code int) string {
	msg, ok := LevelName[code]
	if ok {
		return msg
	}
	return ""
}
