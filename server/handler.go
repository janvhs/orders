package server

import (
	"fmt"
	"net/http"

	"git.bode.fun/orders/order"
)

func (s *server) registerHandlers() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})
	s.router.Route("/order", order.RegisterHandlers(s.db))
}
