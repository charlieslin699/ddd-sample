package repository

import (
	"context"
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/core/repository"
)

// AccountPermissionRepository 帳號權限
type AccountPermissionRepository interface {
	Find(ctx context.Context, accountUID string) (*aggregate.AccountPermission, error)
	Update(ctx context.Context, ap *aggregate.AccountPermission) error

	repository.CoreRepository
}
