package localtime

import (
	"time"
)

type LocalTime interface {
	Location() time.Location
	NowTime() time.Time
	NowTimeString() string
}

type localTime struct {
	location *time.Location
}

// NewLocalTime 建立 LocalTime(預設美東時區)
func NewLocalTime(fns ...localTimeOption) LocalTime {
	location := time.FixedZone("UTC-4", -4*60*60) // 美東時間
	l := &localTime{
		location: location,
	}

	for _, fn := range fns {
		fn(l)
	}

	return l
}

func (l *localTime) Location() time.Location {
	return *l.location
}

func (l *localTime) NowTime() time.Time {
	return time.Now().In(l.location)
}

func (l *localTime) NowTimeString() string {
	return DateTimeString(l.NowTime())
}
