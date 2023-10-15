package util

import (
	"time"
)

var (
	LocalDateTimeFormat string = "2006-01-02 15:04:05"
)

func GetTodayZero() string {

	todayTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Format(LocalDateTimeFormat)
	return todayTime
}

func GetNextTodayZero() string {

	nextTodayTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 24, 0, 0, 0, time.Local).Format(LocalDateTimeFormat)
	return nextTodayTime
}
