package mysql

import (
	"context"
	"ddd-sample/infra/db"
	"ddd-sample/infra/db/auth/model"
)

type ThirdPartyVerification struct {
	conn db.DBConn
}

func NewThirdPartyVerification(conn db.DBConn) *ThirdPartyVerification {
	return &ThirdPartyVerification{
		conn: conn,
	}
}

func (m *ThirdPartyVerification) AddThirdPartyVerification(
	ctx context.Context, thirdPartyVerification model.ThirdPartyVerification,
) error {
	result := m.conn.DB(ctx).Create(&thirdPartyVerification)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *ThirdPartyVerification) GetAccountVerification(
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
