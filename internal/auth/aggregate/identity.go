package aggregate

import (
	"ddd-sample/internal/auth/entity"
	"ddd-sample/internal/auth/event"
	"ddd-sample/internal/auth/valueobject"
	coreaggregate "ddd-sample/internal/core/aggregate"
	"ddd-sample/pkg/token"
	"time"
)

// Identity 身份驗證
type Identity struct {
	account        entity.Account               // 會員資料
	loginFailCount valueobject.LoginFailedCount // 登入失敗次數

	coreaggregate.CoreAggregate
}

// NewIdenetity 建立身份驗證
func NewIdenetity(username, password string) *Identity {
	// 預設值
	i := &Identity{
		account: entity.NewAccount(username, password),

		CoreAggregate: coreaggregate.NewCoreAggregate(),
	}

	return i
}

// BuildIdenetity 建立身份驗證
func BuildIdenetity(account entity.Account) *Identity {
	return &Identity{
		account:       account,
		CoreAggregate: coreaggregate.NewCoreAggregate(),
	}
}

func (i *Identity) Account() entity.Account {
	return i.account
}

// CheckIdentity 檢查帳號密碼
func (i *Identity) CheckIdentity(username, password string) bool {
	isOk := i.account.CheckIdentity(username, password)
	if !isOk {
		i.loginFailCount.Increase()
	}
	return isOk
}

// CreateToken 產生token, 等於登入
func (i *Identity) CreateToken(secretKey []byte, nowTime time.Time) valueobject.Token {
	expirationTime := nowTime.Add(24 * time.Hour) // 24小時到期
	// 產生token
	tokenString := token.NewAuthToken(i.account.UID(), i.account.Username(), secretKey, expirationTime, nowTime)

	// 組資料
	tokenStruct := valueobject.Token{
		TokenString: tokenString,
		CreateTime:  nowTime,
	}

	// 產生事件
	i.AddEvent(event.NewLoginEvent(nowTime, i.account.UID(), i.account.Username()))

	return tokenStruct
}
