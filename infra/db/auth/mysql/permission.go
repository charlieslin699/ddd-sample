package mysql

import (
	"context"
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
)

type Permission struct {
	conn db.DBConn
}

func NewPermission(conn db.DBConn) *Permission {
	return &Permission{
		conn: conn,
	}
}

func (m *Permission) GetAllPermission(ctx context.Context) ([]model.Permission, error) {
	permissions := []model.Permission{}
	result := m.conn.DB(ctx).Find(&permissions)
	if result.Error != nil {
		return nil, result.Error
	}

	return permissions, nil
}
