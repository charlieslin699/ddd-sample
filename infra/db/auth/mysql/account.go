package mysql

import (
	"context"
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
)

type MySQLAccount struct {
	conn db.DBConn
}

func NewMySQLAccount(conn db.DBConn) *MySQLAccount {
	return &MySQLAccount{
		conn: conn,
	}
}

func (m *MySQLAccount) GetAccount(ctx context.Context, uid string) (model.Account, error) {
	accountTable := model.Account{}
	result := m.conn.DB(ctx).Where("uid = ?", uid).
		First(&accountTable)

	if result.Error != nil {
		return model.Account{}, result.Error
	}

	return accountTable, nil
}

func (m *MySQLAccount) GetAccountByUsername(ctx context.Context, username string) (model.Account, error) {
	accountTable := model.Account{}
	result := m.conn.DB(ctx).Where("username = ?", username).
		First(&accountTable)

	if result.Error != nil {
		return model.Account{}, result.Error
	}

	return accountTable, nil
}

func (m *MySQLAccount) AddAccount(ctx context.Context, account model.Account) error {
	result := m.conn.DB(ctx).Create(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MySQLAccount) UpdateAccount(ctx context.Context, account model.Account) error {
	result := m.conn.DB(ctx).
		Where("uid = ?", account.UID).
		Save(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MySQLAccount) ChangePassword(ctx context.Context, uid, password string) error {
	result := m.conn.DB(ctx).
		Model(&model.Account{}).
		Where("uid = ?", uid).
		Update("password", password)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
