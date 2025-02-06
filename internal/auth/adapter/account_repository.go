package adapter

import (
	"context"
	infradbauth "ddd-sample/infra/db/auth"
	infradbauthmodel "ddd-sample/infra/db/auth/model"
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/auth/entity"
	"ddd-sample/internal/auth/repository"
	"ddd-sample/internal/auth/valueobject"
	coreadapter "ddd-sample/internal/core/adapter"
	corerepository "ddd-sample/internal/core/repository"
	"time"
)

type accountRepository struct {
	dbAuth infradbauth.DBAuth

	corerepository.CoreRepository
}

// NewAccountRepository 建立帳號 repository
func NewAccountRepository(dbAuth infradbauth.DBAuth) repository.AccountRepository {
	return &accountRepository{
		dbAuth:         dbAuth,
		CoreRepository: coreadapter.NewCoreRepository(),
	}
}

// New 新增aggregate
func (r *accountRepository) New(username, password string, nowTime time.Time) *aggregate.Account {
	account := aggregate.NewAccount(username, password, nowTime)
	return account
}

// Find 取aggregate
func (r *accountRepository) Find(ctx context.Context, uid string) (*aggregate.Account, error) {
	// 取資料
	accountData, err := r.dbAuth.GetAccount(ctx, uid)
	if err != nil {
		return nil, err
	}

	// 取得帳號的驗證資料
	verifications, err := r.dbAuth.GetAccountVerification(ctx, accountData.UID)
	if err != nil {
		return nil, err
	}

	// 轉換value object
	thirdPartyVerification := r.parseToThirdPartyVerification(verifications)

	account := aggregate.BuildAccount(
		entity.BuildAccount(
			accountData.UID,
			accountData.Username,
			accountData.Password,
			accountData.Secret,
			accountData.Status,
		),
		thirdPartyVerification,
	)

	return account, nil
}

// Add 新增帳號
func (r *accountRepository) Add(ctx context.Context, account *aggregate.Account) error {
	err := r.dbAuth.AddAccount(
		ctx,
		infradbauthmodel.Account{
			UID:      account.Account().UID(),
			Username: account.Account().Username(),
			Password: account.Account().Password(),
			Status:   account.Account().Status(),
			Secret:   account.Account().Secret(),
		},
	)
	if err != nil {
		return err
	}

	return r.PubEvent(account)
}

// Update 更新帳號
func (r *accountRepository) Update(ctx context.Context, account *aggregate.Account) error {
	err := r.dbAuth.UpdateAccount(
		ctx,
		infradbauthmodel.Account{
			UID:    account.Account().UID(),
			Status: account.Account().Status(),
		},
	)
	if err != nil {
		return err
	}

	return r.PubEvent(account)
}

// ChangePassword 更改密碼
func (r *accountRepository) ChangePassword(ctx context.Context, account *aggregate.Account) error {
	err := r.dbAuth.ChangePassword(ctx, account.Account().UID(), account.Account().Password())
	if err != nil {
		return err
	}

	return r.PubEvent(account)
}

// 將第三方驗證資料轉換成value object
func (r *accountRepository) parseToThirdPartyVerification(
	verifications []infradbauthmodel.ThirdPartyVerification,
) valueobject.ThirdPartyVerification {
	thirdPartyVerification := valueobject.NewThirdPartyVerification()

	for _, verification := range verifications {
		thirdPartyVerification.Enable(verification.VerificationType)
	}

	return thirdPartyVerification
}
