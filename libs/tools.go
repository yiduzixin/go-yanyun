package libs

import "time"

func GetDateFormat() string {
	return time.Now().Format("2006/01/02")
}
