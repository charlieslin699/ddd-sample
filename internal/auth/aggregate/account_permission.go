package aggregate

import (
	"ddd-sample/internal/auth/entity"
	"ddd-sample/internal/auth/event"
	"ddd-sample/internal/auth/valueobject"
	"ddd-sample/internal/core/aggregate"
	"time"
)

type AccountPermission struct {
	account     entity.Account
	permissions valueobject.Permissions

	aggregate.CoreAggregate
}

// BuildAccountPermission 建立帳號權限
func BuildAccountPermission(account entity.Account, permissions valueobject.Permissions) *AccountPermission {
	return &AccountPermission{
		account:     account,
		permissions: permissions,

		CoreAggregate: aggregate.NewCoreAggregate(),
	}
}

// Account 取得帳號
func (ap *AccountPermission) Account() entity.Account {
	return ap.account
}

// Permissions 取得權限
func (ap *AccountPermission) Permissions() valueobject.Permissions {
	return ap.permissions
}

// Update 更新權限
func (ap *AccountPermission) Update(newPermissionUIDS []string, nowTime time.Time) {
	ap.permissions.Update(newPermissionUIDS)
	ap.AddEvent(event.NewUpdateAccountPermissionEvent(ap.account.UID(), nowTime))
}
