package entity

import (
	"crypto/sha256"
	"ddd-sample/internal/auth/enum"
	"ddd-sample/internal/core/entity"
	"ddd-sample/pkg/uid"
	"encoding/hex"
)

type Account struct {
	username string // 帳號名稱(禁止變更)
	password string
	secret   string
	status   enum.AccountStatus

	entity.CoreEntity
}

// NewAccount 建立帳號
func NewAccount(username, password string, fns ...AccountOptionFunc) Account {
	a := Account{
		username:   username,
		secret:     uid.NewNanoID(),          // 產生帳號金鑰
		status:     enum.AccountStatusNormal, // 預設狀態為正常
		CoreEntity: entity.NewCoreEntity(),
	}

	a.password = a.cryptoPassword(password)

	for _, fn := range fns {
		fn(&a)
	}

	return a
}

// BuildAccount 建立帳號
func BuildAccount(uid, username, password, secret string, status enum.AccountStatus) Account {
	return Account{
		username:   username,
		password:   password,
		secret:     secret,
		status:     status,
		CoreEntity: entity.BuildCoreEntity(uid),
	}
}

// Username 取得帳號名稱
func (a Account) Username() string {
	return a.username
}

// Password 取得密碼
func (a Account) Password() string {
	return a.password
}

// Secret 取得帳號金鑰
func (a Account) Secret() string {
	return a.secret
}

// Status 取得帳號狀態
func (a Account) Status() enum.AccountStatus {
	return a.status
}

// Update 更新帳號
func (a *Account) Update(status enum.AccountStatus) {
	a.status = status
}

// ChangePassword 更改密碼
func (a *Account) ChangePassword(password string) {
	a.password = a.cryptoPassword(password)
}

// 檢查帳號密碼
func (a *Account) CheckIdentity(username, password string) bool {
	cryptoPassword := a.cryptoPassword(password)

	return a.username == username && a.password == cryptoPassword
}

// 加密密碼
func (a *Account) cryptoPassword(password string) string {
	sha256Bytes := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sha256Bytes[:])
}

// AccountOptionFunc 帳號選項
type AccountOptionFunc func(a *Account)

// AccountWithUID 帳號UID
func AccountWithSecret(secret string) AccountOptionFunc {
	return func(a *Account) {
		a.secret = secret
	}
}

// AccountWithStatus 帳號狀態
func AccountWithStatus(status enum.AccountStatus) AccountOptionFunc {
	return func(a *Account) {
		a.status = status
	}
}
