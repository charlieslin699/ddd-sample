package model

type AccountPermission struct {
	ID            int64  `gorm:"column:id;primaryKey"`
	AccountUID    string `gorm:"column:accountUID"`
	PermissionUID string `gorm:"column:permissionUID"`
}

func (AccountPermission) TableName() string {
	return "accountPermission"
}
