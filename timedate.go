package goUtils

import (
	"strconv"
	"strings"
	"time"
)

// GetTodayYMD 得到以sep为分隔符的年、月、日字符串(今天)
func GetTodayYMD(sep string) string {
	now := time.Now()
	todayYMD := now.Format("2006-01-02")

	return strings.Replace(todayYMD, "-", sep, -1)
}

// GetAnydayYMD 得到以sep为分隔符的年、月、日字符串(以今天为准的任何一天)
func GetAnydayYMD(sep string, any int) string {
	if any == 0 {
		return GetTodayYMD(sep)
	} else {
		now := time.Now()
		changeStr := strconv.Itoa(any*24) + "h"
		changeTime, _ := time.ParseDuration(changeStr)
		anyday := now.Add(changeTime)
		anydayYMD := anyday.Format("2006-01-02")

		return strings.Replace(anydayYMD, "-", sep, -1)
	}
}

// GetTodayYM 得到以sep为分隔符的年、月字符串(今天)
func GetTodayYM(sep string) string {
	now := time.Now()
	todayYM := now.Format("2006-01")

	return strings.Replace(todayYM, "-", sep, -1)
}
