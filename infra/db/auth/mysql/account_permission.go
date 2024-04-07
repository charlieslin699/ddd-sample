package mysql

import (
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
)

type MySQLAccountPermission struct {
	conn db.DBConn
}

func NewMySQLAccountPermission(conn db.DBConn) *MySQLAccountPermission {
	return &MySQLAccountPermission{
		conn: conn,
	}
}

// 取得帳號權限
func (m *MySQLAccountPermission) GetAccountPermission(accountUID string) ([]model.AccountPermission, error) {
	accountPermissions := []model.AccountPermission{}
	result := m.conn.DB().Where("accountUID = ?", accountUID).
		Find(&accountPermissions)

	if result.Error != nil {
		return nil, result.Error
	}

	return accountPermissions, nil
}

// 更新帳號權限(全部刪除後寫入)
func (m *MySQLAccountPermission) UpdateAccountPermission(accountUID string, permissions []model.AccountPermission) error {
	result := m.conn.DB().Where("accountUID = ?", accountUID).Delete(model.AccountPermission{})
	if result.Error != nil {
		return result.Error
	}

	result = m.conn.DB().Create(permissions)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
