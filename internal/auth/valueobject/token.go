package valueobject

import "time"

type Token struct {
	TokenString string
	CreateTime  time.Time
}
