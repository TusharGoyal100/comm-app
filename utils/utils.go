package utils

import "time"

func CreateTimeStamp() string {
	t := time.Now()
	return t.Format("20060102150405")
}
