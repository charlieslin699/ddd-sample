package event

import (
	"ddd-sample/internal/core/event"
	"ddd-sample/pkg/localtime"
	"time"
)

// UpdateAccountEvent 更新帳號事件
type UpdateAccountEvent struct {
	event.CoreEvent[UpdateAccountEventData]
}

// UpdateAccountEventData 更新帳號事件資料
type UpdateAccountEventData struct {
	UID string `json:"uid"`
}

// NewUpdateAccountEvent 建立更新帳號事件
func NewUpdateAccountEvent(nowTime time.Time, uid string) event.Event {
	e := UpdateAccountEvent{}
	e.CoreEvent = event.NewCoreEvent(
		localtime.DateTimeString(nowTime),
		UpdateAccountEventData{
			UID: uid,
		},
	)

	return e
}
