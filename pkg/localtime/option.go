package localtime

import "time"

type localTimeOption func(l *localTime)

func WithLocation(location time.Location) localTimeOption {
	return func(l *localTime) {
		l.location = &location
	}
}
