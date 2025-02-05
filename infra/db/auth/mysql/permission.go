package mysql

import (
	"context"
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
)

type MySQLPermission struct {
	conn db.DBConn
}

func NewMySQLPermission(conn db.DBConn) *MySQLPermission {
	return &MySQLPermission{
		conn: conn,
	}
}

func (m *MySQLPermission) GetAllPermission(ctx context.Context) ([]model.Permission, error) {
	permissions := []model.Permission{}
	result := m.conn.DB(ctx).Find(&permissions)
	if result.Error != nil {
		return nil, result.Error
	}

	return permissions, nil
}
