package model

import (
	"time"
)

type LoginRecord struct {
	ID         int64     `gorm:"column:id;primaryKey"`
	AccountUID string    `gorm:"column:accountUID"`
	Token      string    `gorm:"column:token"`      // 登入token
	CreatedAt  time.Time `gorm:"column:createTime"` // 登入時間, gorm自動帶入
}

func (LoginRecord) TableName() string {
	return "loginRecord"
}
