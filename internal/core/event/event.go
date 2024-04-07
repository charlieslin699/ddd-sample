package event

import "ddd-sample/pkg/errorcode"

type Event interface {
	GetName() string
	ParseToJSON() (string, errorcode.ErrorCode)
}
