package mysql

import (
	"context"
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
)

type MySQLThirdPartyVerification struct {
	conn db.DBConn
}

func NewMySQLThirdPartyVerification(conn db.DBConn) *MySQLThirdPartyVerification {
	return &MySQLThirdPartyVerification{
		conn: conn,
	}
}

func (m *MySQLThirdPartyVerification) AddThirdPartyVerification(
	ctx context.Context, thirdPartyVerification model.ThirdPartyVerification,
) error {
	result := m.conn.DB(ctx).Create(&thirdPartyVerification)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MySQLThirdPartyVerification) GetAccountVerification(
	ctx context.Context, accountUID string,
) ([]model.ThirdPartyVerification, error) {
	thirdPartyVerification := []model.ThirdPartyVerification{}
	result := m.conn.DB(ctx).Where("accountUID = ?", accountUID).
		Find(&thirdPartyVerification)

	if result.Error != nil {
		return nil, result.Error
	}

	return thirdPartyVerification, nil
}
