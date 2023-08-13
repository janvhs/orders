package order

import (
	"database/sql"
	"time"
)

type State = int64

const (
	Open     State = 0
	Ordered  State = 1
	Rejected State = 2
)

type Order struct {
	ID int64

	State State
	// TODO: Move to article?
	Until     sql.NullTime
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`

	AuthorID int64 `db:"author_id"`
}
