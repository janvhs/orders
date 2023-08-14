package entity

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	// Primary Key
	ID int64 `bun:",pk,autoincrement"`

	// Attributes
	GUID     string
	Username string
	UserDN   string

	// Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `bun:",soft_delete,nullzero"`
}
