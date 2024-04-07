package log

import "time"

var timeLocation *time.Location // 時區
var timeFormat string           // 時間格式

func init() {
	timeLocation = time.FixedZone("UTC+8", 8*60*60) // 台灣時間
	timeFormat = time.DateTime
}

func SetTimeLocation(location time.Location) {
	timeLocation = &location
}

func SetTimeFormat(format string) {
	timeFormat = format
}

func nowTime() string {
	return "[" + time.Now().In(timeLocation).Format(timeFormat) + "]"
}
