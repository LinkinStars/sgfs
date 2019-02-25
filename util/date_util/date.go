package date_util

import (
	"time"
)

const (
	YYYYMMDD       = "2006-01-02"
	YYYYMMddHHmmss = "20060102150405"
)

func GetCurTimeFormat(format string) string {
	cur := time.Now()
	return cur.Format(format)
}
