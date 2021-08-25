package times

import (
	"errors"
	"strings"
	"time"
)

// Str2TimeStruct 将字符串转换为时间结构
func Str2TimeStruct(str string, format ...string) (time.Time, error) {
	f := ""
	if len(format) > 0 {
		f = strings.Trim(format[0], " ")
	} else {
		f = "2006-01-02 15:04:05"
	}
	if len(str) != len(f) {
		return time.Now(), errors.New("parameter format error")
	}
	return time.Parse(f, str)
}