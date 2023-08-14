package entity

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type OrderState = int64

const (
	OrderOpen     OrderState = 0
	OrderOrdered  OrderState = 1
	OrderRejected OrderState = 2
)

type Order struct {
	bun.BaseModel `bun:"table:orders"`

	// Primary Key
	ID int64 `bun:",pk,autoincrement"`

	// Attributes
	State OrderState

	// Foreign Keys
	AuthorID int64

	// ORM-Relationships
	// TODO: Replace pointer with sql.Null like struct. Could be a generic.
	Author *User `bun:"rel:belongs-to"`
	// Articles []*Article `bun:"rel:has-many"`

	// Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `bun:",soft_delete,nullzero"`
}
