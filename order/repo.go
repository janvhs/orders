package order

import (
	"github.com/jmoiron/sqlx"
)

// Interface Abstraction
// ------------------------------------------------------------------------

type Repo interface {
	GetAll() ([]Order, error)
}

// Main implementation
// ------------------------------------------------------------------------

var _ Repo = (*sqlRepo)(nil)

type sqlRepo struct {
	db sqlx.DB
}

func NewRepo(db *sqlx.DB) *sqlRepo {
	return &sqlRepo{
		db: *db,
	}
}

// Public Methods
// ------------------------------------------------------------------------

func (r *sqlRepo) GetAll() ([]Order, error) {
	// FIXME: Return an actual value
	var orders []Order
	err := r.db.Select(&orders, "SELECT * FROM orders")
	return orders, err
}
