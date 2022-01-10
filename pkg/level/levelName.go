package level

var LevelName = map[int]string{
	EASY:      "简单",
	MIDDLE:    "中等",
	HARD:      "困难",
	HARD_MORE: "极其困难",
}

func GetLevelName(code int) string {
	msg, ok := LevelName[code]
	if ok {
		return msg
	}
	return ""
}
