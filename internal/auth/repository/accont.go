package repository

import (
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/core/repository"
	"time"
)

// AccountRepository 帳號
type AccountRepository interface {
	New(username, password string, nowTime time.Time) *aggregate.Account
	Find(uid string) (*aggregate.Account, error)
	Add(account *aggregate.Account) error
	Update(account *aggregate.Account) error
	ChangePassword(account *aggregate.Account) error

	repository.CoreRepository
}
