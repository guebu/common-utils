package date_utils

import (
	"time"
)

const (
	DateFormatString = "2006-01-02T15:04:05Z"
)

func GetNowString() string {
	now := GetNow()
	return now.Format(DateFormatString)
}

func GetNow() time.Time {
	return time.Now().UTC()
}
