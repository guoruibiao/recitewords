package utils

import (
	"time"
	"strings"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

// 获取今天格式化的日期串
func GetCurDate() string {
	return time.Now().Format("2006-01-02")
}

// 仅提供 YYYY-MM-DD 和 YYYY-MM-DD HH:II:SS 格式
func DateString(format string, timestamp int) (date string) {
	var layout string
	if strings.Contains(format, " ") == false {
		layout = strings.Split(TIME_LAYOUT, " ")[0]
	}else {
		layout = TIME_LAYOUT
	}
	date = time.Unix(int64(timestamp), 0).Format(layout)
	return
}

// 将时间转成 unix 时间戳
func StringToTime(date string) (timestamp int) {
	if date == "" {
		return int(time.Now().Unix())
	}
	if timeObject, err := time.Parse(TIME_LAYOUT, date); err != nil {
		return
	}else{
		timestamp = int(timeObject.Unix())
	}
	return
}

