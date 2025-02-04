package mysql

import (
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
	"ddd-sample/pkg/errorcode"
	"errors"

	"gorm.io/gorm"
)

type MySQLAccount struct {
	conn db.DBConn
}

func NewMySQLAccount(conn db.DBConn) *MySQLAccount {
	return &MySQLAccount{
		conn: conn,
	}
}

func (m *MySQLAccount) GetAccount(uid string) (model.Account, error) {
	accountTable := model.Account{}
	result := m.conn.DB().Where("uid = ?", uid).
		First(&accountTable)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Account{}, errorcode.ErrDBRecordNotFound
		}
		return model.Account{}, result.Error
	}

	return accountTable, nil
}

func (m *MySQLAccount) GetAccountByUsername(username string) (model.Account, error) {
	accountTable := model.Account{}
	result := m.conn.DB().Where("username = ?", username).
		First(&accountTable)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Account{}, errorcode.ErrDBRecordNotFound
		}
		return model.Account{}, result.Error
	}

	return accountTable, nil
}

func (m *MySQLAccount) AddAccount(account model.Account) error {
	result := m.conn.DB().Create(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MySQLAccount) UpdateAccount(account model.Account) error {
	result := m.conn.DB().Where("uid = ?", account.UID).Save(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MySQLAccount) ChangePassword(uid, password string) error {
	result := m.conn.DB().Model(&model.Account{}).Where("uid = ?", uid).Update("password", password)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
