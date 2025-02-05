package repository

import (
	"context"
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/core/repository"
	"time"
)

// AccountRepository 帳號
type AccountRepository interface {
	New(username, password string, nowTime time.Time) *aggregate.Account
	Find(ctx context.Context, uid string) (*aggregate.Account, error)
	Add(ctx context.Context, account *aggregate.Account) error
	Update(ctx context.Context, account *aggregate.Account) error
	ChangePassword(ctx context.Context, account *aggregate.Account) error

	repository.CoreRepository
}
