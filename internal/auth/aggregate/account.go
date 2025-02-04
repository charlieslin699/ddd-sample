package aggregate

import (
	"ddd-sample/internal/auth/entity"
	"ddd-sample/internal/auth/enum"
	"ddd-sample/internal/auth/event"
	"ddd-sample/internal/auth/valueobject"
	"ddd-sample/internal/core/aggregate"
	"time"
)

// Account 帳號
type Account struct {
	account                entity.Account
	thirdPartyVerification valueobject.ThirdPartyVerification

	aggregate.CoreAggregate
}

// NewAccount 建立帳號
func NewAccount(username, password string, nowTime time.Time) *Account {
	a := &Account{
		account: entity.NewAccount(username, password),

		CoreAggregate: aggregate.NewCoreAggregate(),
	}

	a.AddEvent(event.NewCreateAccountEvent(nowTime, a.account.UID(), a.account.Username()))

	return a
}

// BuildAccount 建立帳號
func BuildAccount(account entity.Account, thirdPartyVerification valueobject.ThirdPartyVerification) *Account {
	return &Account{
		account:                account,
		thirdPartyVerification: thirdPartyVerification,
		CoreAggregate:          aggregate.NewCoreAggregate(),
	}
}

// Account 取得帳號
func (a *Account) Account() entity.Account {
	return a.account
}

// ThirdPartyVerification 取得第三方驗證
func (a *Account) ThirdPartyVerification() valueobject.ThirdPartyVerification {
	return a.thirdPartyVerification
}

// Update 更新帳號
func (a *Account) Update(status enum.AccountStatus, isEnabledOTP bool, nowTime time.Time) {
	a.account.Update(status)
	a.thirdPartyVerification.OTP = isEnabledOTP
	a.AddEvent(event.NewUpdateAccountEvent(nowTime, a.account.UID()))
}

// ChangePassword 更改密碼
func (a *Account) ChangePassword(password string, nowTime time.Time) {
	a.account.ChangePassword(password)
	a.AddEvent(event.NewChangePasswordEvent(nowTime, a.account.UID()))
}

// PasswordVerify 密碼驗證
func (a *Account) PasswordVerify(password string) bool {
	return a.account.CheckIdentity(a.account.Username(), password)
}
