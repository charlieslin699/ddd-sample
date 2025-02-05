package mysql

import (
	"context"
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
)

type MySQLLoginRecord struct {
	conn db.DBConn
}

func NewMySQLLoginRecord(conn db.DBConn) *MySQLLoginRecord {
	return &MySQLLoginRecord{
		conn: conn,
	}
}

func (m *MySQLLoginRecord) AddLoginRecord(ctx context.Context, accountUID, token string) error {
	recordTable := model.LoginRecord{
		AccountUID: accountUID,
		Token:      token,
	}

	result := m.conn.DB(ctx).Create(&recordTable)
	return result.Error
}
