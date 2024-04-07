package model

import "time"

type Account struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	UID       string    `gorm:"column:uid"`
	Username  string    `gorm:"column:username"`   // 帳號
	Password  string    `gorm:"column:password"`   // 密碼
	Status    uint      `gorm:"column:status"`     // 帳號狀態
	Secret    string    `gorm:"column:secret"`     // 帳號金鑰
	CreatedAt time.Time `gorm:"column:createTime"` // gorm自動帶入
	UpdatedAt time.Time `gorm:"column:updateTime"` // gorm自動帶入
}

func (Account) TableName() string {
	return "account"
}
