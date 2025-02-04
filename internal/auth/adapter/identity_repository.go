package adapter

import (
	infradbauth "ddd-sample/infra/db/auth"
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/auth/entity"
	"ddd-sample/internal/auth/enum"
	"ddd-sample/internal/auth/repository"
	"ddd-sample/internal/auth/valueobject"
	coreradapter "ddd-sample/internal/core/adapter"
	corerepository "ddd-sample/internal/core/repository"
)

type identityRepository struct {
	dbAuth infradbauth.DBAuth

	corerepository.CoreRepository
}

// NewIdentityRepository 建立身份驗證 repository
func NewIdentityRepository(mysqlAuth infradbauth.DBAuth) repository.IdentityRepository {
	return &identityRepository{
		dbAuth:         mysqlAuth,
		CoreRepository: coreradapter.NewCoreRepository(),
	}
}

// Find 取aggregate
func (repo *identityRepository) Find(username string) (*aggregate.Identity, error) {
	// 取資料
	accountTable, err := repo.dbAuth.GetAccountByUsername(username)
	if err != nil {
		return nil, err
	}

	// 轉換列舉
	status, err := enum.ConvertToAccountStatus(accountTable.Status)
	if err != nil {
		return nil, err
	}

	identity := aggregate.BuildIdenetity(
		entity.BuildAccount(
			accountTable.UID,
			accountTable.Username,
			accountTable.Password,
			accountTable.Secret,
			status,
		),
	)

	return identity, nil
}

// SaveLoginRecord 登入紀錄
func (repo *identityRepository) SaveLoginRecord(identity *aggregate.Identity, token valueobject.Token) error {
	if err := repo.dbAuth.AddLoginRecord(
		identity.Account().UID(),
		token.TokenString,
	); err != nil {
		return err
	}

	return repo.PubEvent(identity)
}

// SaveLoginFailedRecord 登入失敗紀錄
func (repo *identityRepository) SaveLoginFailedRecord(identity *aggregate.Identity) error {
	// TODO: Redis寫入失敗次數

	return repo.PubEvent(identity)
}
