package event

import (
	"ddd-sample/internal/core/event"
	"ddd-sample/pkg/localtime"
	"time"
)

// CreateAccountEvent 建立帳號事件
type CreateAccountEvent struct {
	event.CoreEvent[CreateAccountEventData]
}

// CreateAccountEventData 建立帳號事件資料
type CreateAccountEventData struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
}

// NewCreateAccountEvent 建立建立帳號事件
func NewCreateAccountEvent(nowTime time.Time, uid, username string) event.Event {
	e := CreateAccountEvent{}
	e.CoreEvent = event.NewCoreEvent(
		localtime.DateTimeString(nowTime),
		CreateAccountEventData{
			UID:      uid,
			Username: username,
		},
	)

	return e
}
