package event

import (
	"ddd-sample/internal/core/event"
	"ddd-sample/pkg/localtime"
	"time"
)

// UpdateAccountPermissionEvent 更新帳號權限事件
type UpdateAccountPermissionEvent struct {
	event.CoreEvent[UpdateAccountPermissionEventData]
}

// UpdateAccountPermissionEventData 更新帳號權限事件資料
type UpdateAccountPermissionEventData struct {
	AccountUID string `json:"account_uid"`
}

// NewUpdateAccountPermissionEvent 建立更新帳號權限事件
func NewUpdateAccountPermissionEvent(accountUID string, nowTime time.Time) event.Event {
	e := UpdateAccountPermissionEvent{}
	e.CoreEvent = event.NewCoreEvent(
		localtime.DateTimeString(nowTime),
		UpdateAccountPermissionEventData{
			AccountUID: accountUID,
		},
	)

	return e
}
