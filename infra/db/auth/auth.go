// 紀錄Auth DB相關資料
package auth

import (
	"context"
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
	"ddd-sample/infra/db/auth/mysql"
)

// 資料庫 Auth
type DBAuth interface {
	TableAccount
	TableAccountPermission
	TablePermission
	TableLoginRecord
	TableThirdPartyVerification
}

// 資料表 Account
type TableAccount interface {
	GetAccount(ctx context.Context, uid string) (model.Account, error)
	GetAccountByUsername(ctx context.Context, username string) (model.Account, error)
	AddAccount(ctx context.Context, account model.Account) error
	UpdateAccount(ctx context.Context, account model.Account) error
	ChangePassword(ctx context.Context, uid, password string) error
}

// 資料表 AccountPermission
type TableAccountPermission interface {
	GetAccountPermission(ctx context.Context, accountUID string) ([]model.AccountPermission, error)
	UpdateAccountPermission(ctx context.Context, accountUID string, permissions []model.AccountPermission) error // 更新帳號權限(全部刪除後寫入)
}

// 資料表 Permission
type TablePermission interface {
	GetAllPermission(ctx context.Context) ([]model.Permission, error)
}

// 資料表 LoginRecord
type TableLoginRecord interface {
	AddLoginRecord(ctx context.Context, username, token string) error
}

// 資料表 ThirdPartyVerification
type TableThirdPartyVerification interface {
	AddThirdPartyVerification(ctx context.Context, tpv model.ThirdPartyVerification) error
	GetAccountVerification(ctx context.Context, accountUID string) ([]model.ThirdPartyVerification, error)
}

// MySQL工廠
func NewMySQLAuth(conn db.DBConn) DBAuth {
	return &struct {
		TableAccount
		TableAccountPermission
		TablePermission
		TableLoginRecord
		TableThirdPartyVerification
	}{
		TableAccount:                mysql.NewMySQLAccount(conn),
		TableAccountPermission:      mysql.NewMySQLAccountPermission(conn),
		TableLoginRecord:            mysql.NewMySQLLoginRecord(conn),
		TablePermission:             mysql.NewMySQLPermission(conn),
		TableThirdPartyVerification: mysql.NewMySQLThirdPartyVerification(conn),
	}
}
