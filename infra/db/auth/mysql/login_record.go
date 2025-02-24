package mysql

import (
	"context"
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
)

type LoginRecord struct {
	conn db.DBConn
}

func NewLoginRecord(conn db.DBConn) *LoginRecord {
	return &LoginRecord{
		conn: conn,
	}
}

func (m *LoginRecord) AddLoginRecord(ctx context.Context, accountUID, token string) error {
	recordTable := model.LoginRecord{
		AccountUID: accountUID,
		Token:      token,
	}

	result := m.conn.DB(ctx).Create(&recordTable)
	return result.Error
}
