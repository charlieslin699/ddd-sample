package log

import "time"

var timeLocation = time.FixedZone("UTC+8", 8*60*60) //nolint:mnd // 時區
var timeFormat = time.DateTime                      // 時間格式

func SetTimeLocation(location time.Location) {
	timeLocation = &location
}

func SetTimeFormat(format string) {
	timeFormat = format
}

func nowTime() string {
	return "[" + time.Now().In(timeLocation).Format(timeFormat) + "]"
}
