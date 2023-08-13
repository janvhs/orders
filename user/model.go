package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID   int64
	GUID string

	Username string
	UserDN   string `db:"user_dn"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
