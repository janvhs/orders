package server

import (
	"fmt"
	"net/http"
)

func (s *server) registerHandlers() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})
}
