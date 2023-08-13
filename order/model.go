package order

import (
	"time"

	"git.bode.fun/orders/article"
)

type State = int

const (
	Open State = iota
	Ordered
	Rejected
)

type Order struct {
	ID int64

	State     State
	Until     time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// TODO: Replace with User
	Author   string
	Articles []article.Article
}
