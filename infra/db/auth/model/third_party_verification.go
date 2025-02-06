package model

import (
	"ddd-sample/internal/auth/enum"
	"time"
)

type ThirdPartyVerification struct {
	ID               int64                       `gorm:"column:id;primaryKey"`
	AccountUID       string                      `gorm:"column:accountUID"`
	VerificationType enum.ThirdPartyVerification `gorm:"column:verificationType"` // 驗證類型(1: FOTP)
	CreatedAt        time.Time                   `gorm:"column:createTime"`       // gorm自動帶入
}

func (ThirdPartyVerification) TableName() string {
	return "thirdPartyVerification"
}
