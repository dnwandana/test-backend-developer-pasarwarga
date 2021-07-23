package entity

import (
	"database/sql"
	"time"
)

type Category struct {
	ID           int64
	CategoryName string
	CategorySlug string
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
	DeletedAt    sql.NullTime
}
