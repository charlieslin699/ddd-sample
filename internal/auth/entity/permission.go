package entity

import (
	"ddd-sample/internal/core/entity"
)

type Permission struct {
	name  string
	index string

	entity.CoreEntity
}

// NewPermission 建立權限
func BuildPermission(uid, name, key string) Permission {
	return Permission{
		name:       name,
		index:      key,
		CoreEntity: entity.BuildCoreEntity(uid),
	}
}

// Name 取得名稱
func (p Permission) Name() string {
	return p.name
}

// Index 取得索引
func (p Permission) Index() string {
	return p.index
}
