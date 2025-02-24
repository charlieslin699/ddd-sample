package mysql

import (
	"context"
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
)

type Account struct {
	conn db.DBConn
}

func NewAccount(conn db.DBConn) *Account {
	return &Account{
		conn: conn,
	}
}

func (m *Account) GetAccount(ctx context.Context, uid string) (model.Account, error) {
	accountTable := model.Account{}
	result := m.conn.DB(ctx).Where("uid = ?", uid).
		First(&accountTable)

	if result.Error != nil {
		return model.Account{}, result.Error
	}

	return accountTable, nil
}

func (m *Account) GetAccountByUsername(ctx context.Context, username string) (model.Account, error) {
	accountTable := model.Account{}
	result := m.conn.DB(ctx).Where("username = ?", username).
		First(&accountTable)

	if result.Error != nil {
		return model.Account{}, result.Error
	}

	return accountTable, nil
}

func (m *Account) AddAccount(ctx context.Context, account model.Account) error {
	result := m.conn.DB(ctx).Create(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Account) UpdateAccount(ctx context.Context, account model.Account) error {
	result := m.conn.DB(ctx).
		Where("uid = ?", account.UID).
		Save(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Account) ChangePassword(ctx context.Context, uid, password string) error {
	result := m.conn.DB(ctx).
		Model(&model.Account{}).
		Where("uid = ?", uid).
		Update("password", password)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
