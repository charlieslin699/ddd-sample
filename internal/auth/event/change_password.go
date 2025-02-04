package event

import (
	"ddd-sample/internal/core/event"
	"ddd-sample/pkg/localtime"
	"time"
)

// ChangePasswordEvent 更改密碼事件
type ChangePasswordEvent struct {
	event.CoreEvent[ChangePasswordEventData]
}

// ChangePasswordEventData 更改密碼事件資料
type ChangePasswordEventData struct {
	UID string `json:"uid"`
}

// NewChangePasswordEvent 建立更改密碼事件
func NewChangePasswordEvent(nowTime time.Time, uid string) event.Event {
	e := ChangePasswordEvent{}
	e.CoreEvent = event.NewCoreEvent(
		localtime.DateTimeString(nowTime),
		ChangePasswordEventData{
			UID: uid,
		},
	)

	return e
}
