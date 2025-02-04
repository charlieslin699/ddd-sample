package localtime

import "time"

const (
	FormatDate     = time.DateOnly // 僅日期格式 (yyyy-MM-dd)
	FormatDateTime = time.DateTime // 一般日期格式 (yyyy-MM-dd HH:mm:ss)
)

// 僅日期格式 (yyyy-MM-dd)
func DateString(t time.Time) string {
	return t.Format(FormatDate)
}

// 一般日期格式 (yyyy-MM-dd HH:mm:ss)
func DateTimeString(t time.Time) string {
	return t.Format(FormatDateTime)
}
