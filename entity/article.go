package entity

import (
	"database/sql"
	"time"
)

type Article struct {
	ID         int64
	Title      string
	Slug       string
	CategoryID int64
	Content    string
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
	DeletedAt  sql.NullTime
}
