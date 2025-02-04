package repository

import (
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/auth/valueobject"
	"ddd-sample/internal/core/repository"
)

// IdentityRepository 身份驗證
type IdentityRepository interface {
	Find(username string) (*aggregate.Identity, error)                           // 取aggregate
	SaveLoginRecord(identity *aggregate.Identity, token valueobject.Token) error // 登入紀錄
	SaveLoginFailedRecord(identity *aggregate.Identity) error                    // 登入失敗紀錄

	repository.CoreRepository
}
