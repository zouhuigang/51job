package util

import (
	"strconv"
	"time"
)

//今天的日期
func TodayDate() string {
	return time.Now().Format("20060102")
}

//获取时间戳
func TodayTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func SixMonthAgo() string {
	t := time.Now()
	ts := time.Date(t.Year(), t.Month()-6, t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local).Format("20060102")
	return ts
}
