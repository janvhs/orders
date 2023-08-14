package order

import (
	"context"

	"github.com/uptrace/bun"
)

// Interface Abstraction
// ------------------------------------------------------------------------

type Repo interface {
	GetAllRel() ([]Order, error)
}

// Main implementation
// ------------------------------------------------------------------------

var _ Repo = (*sqlRepo)(nil)

type sqlRepo struct {
	db bun.DB
}

func NewRepo(db *bun.DB) *sqlRepo {
	return &sqlRepo{
		db: *db,
	}
}

// Public Methods
// ------------------------------------------------------------------------

func (r *sqlRepo) GetAllRel() ([]Order, error) {
	// FIXME: Return an actual value
	var orders []Order
	err := r.db.NewSelect().
		Model(&orders).
		Relation("Author").
		Scan(context.TODO())

	return orders, err
}
