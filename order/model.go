package order

import "git.bode.fun/orders/db/entity"

type State = entity.OrderState

const (
	Open     State = entity.ArticleOpen
	Ordered  State = entity.ArticleOrdered
	Rejected State = entity.ArticleRejected
)

type Order = entity.Order
