package repository

import (
	"context"
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/auth/valueobject"
	"ddd-sample/internal/core/repository"
)

// IdentityRepository 身份驗證
type IdentityRepository interface {
	Find(ctx context.Context, username string) (*aggregate.Identity, error)                           // 取aggregate
	SaveLoginRecord(ctx context.Context, identity *aggregate.Identity, token valueobject.Token) error // 登入紀錄
	SaveLoginFailedRecord(ctx context.Context, identity *aggregate.Identity) error                    // 登入失敗紀錄

	repository.CoreRepository
}
