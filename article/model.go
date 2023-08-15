package article

import "git.bode.fun/orders/internal/db/entity"

type Article = entity.Article

type State = entity.ArticleState

const (
	Open     State = entity.ArticleOpen
	Ordered  State = entity.ArticleOrdered
	Rejected State = entity.ArticleRejected
)
