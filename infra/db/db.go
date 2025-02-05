package db

import (
	"context"

	"gorm.io/gorm"
)

type DBConn interface {
	DB(ctx context.Context) *gorm.DB
}
