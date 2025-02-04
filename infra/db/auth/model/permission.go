package model

type Permission struct {
	ID        int64  `gorm:"column:id;primaryKey"`
	UID       string `gorm:"column:uid"`
	Group     string `gorm:"column:group"`     // 權限群組
	Name      string `gorm:"column:name"`      // 權限名稱
	Key       string `gorm:"column:key"`       // 權限key
	LangIndex string `gorm:"column:langIndex"` // 字典檔索引
}

func (Permission) TableName() string {
	return "permission"
}
