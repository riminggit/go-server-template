package level

var LevelName = map[string]string{
	
}

func GetLevelName(code string) string {
	msg, ok := LevelName[code]
	if ok {
		return msg
	}
	return ""
}