package entity

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type ArticleState = int64

const (
	ArticleOpen     ArticleState = 0
	ArticleOrdered  ArticleState = 1
	ArticleRejected ArticleState = 2
)

type Article struct {
	bun.BaseModel `bun:"table:articles"`

	// Primary Key
	ID int64 `bun:",pk,autoincrement"`

	// Attributes
	Name       string
	Amount     int64
	URL        string
	Price      float64
	CostCentre string
	State      ArticleState
	Until      sql.NullTime

	// Foreign Keys
	OrderID   int64
	AuthorID  int64
	ForUserID int64

	// ORM-Relationships
	// TODO: Replace pointer with sql.Null like struct. Could be a generic.
	Order   *Order `bun:"rel:belongs-to"`
	Author  *User  `bun:"rel:belongs-to"`
	ForUser *User  `bun:"rel:belongs-to"`

	// Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `bun:",soft_delete,nullzero"`
}
