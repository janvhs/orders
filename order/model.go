package order

import "git.bode.fun/orders/internal/db/entity"

type State = entity.OrderState

const (
	Open     State = entity.ArticleOpen
	Ordered  State = entity.ArticleOrdered
	Rejected State = entity.ArticleRejected
)

type Order = entity.Order
