package repository

import (
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/core/repository"
)

// AccountPermissionRepository 帳號權限
type AccountPermissionRepository interface {
	Find(accountUID string) (*aggregate.AccountPermission, error)
	Update(*aggregate.AccountPermission) error

	repository.CoreRepository
}
