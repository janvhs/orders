package article

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

type Article struct {
	ID int64

	Name       string
	Amount     int64
	URL        string
	Price      float64
	CostCentre string `db:"cost_centre"`
	State      State

	OrderID int64 `db:"order_id"`

	AutorID   int64 `db:"author_id"`
	ForUserID int64 `db:"for_user_id"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
