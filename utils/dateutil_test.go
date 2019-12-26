package utils

import (
	"testing"
	"time"
)

func Test_GetCurDate(t *testing.T) {
	t.Log(GetCurDate())
}


func TestStringToTime(t *testing.T) {
	t.Log(StringToTime(""))
	t.Log(StringToTime("2019-12-26 11:56:00"))
}


func TestDateString(t *testing.T) {
	t.Log(DateString("yyyy-mm-dd", int(time.Now().Unix())))
	t.Log(DateString("yyyy-mm-dd hh:ii:ss", int(time.Now().Unix())))
}