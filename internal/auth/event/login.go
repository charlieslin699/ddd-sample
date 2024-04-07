package event

import (
	"ddd-sample/internal/core/event"
	"ddd-sample/pkg/localtime"
	"time"
)

// LoginEvent 登入事件
type LoginEvent struct {
	event.CoreEvent[LoginEventData]
}

// LoginEventData 登入事件資料
type LoginEventData struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
}

// NewLoginEvent 建立登入事件
func NewLoginEvent(nowTime time.Time, uid, username string) event.Event {
	e := LoginEvent{}
	e.CoreEvent = event.NewCoreEvent(
		localtime.DateTimeString(nowTime),
		LoginEventData{
			UID:      uid,
			Username: username,
		},
	)

	return e
}
