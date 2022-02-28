package util

import "time"

func GetNowTimeUnix() int64 {
	return time.Now().Unix()
}
