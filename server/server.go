package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/uptrace/bun"
)

var _ http.Handler = (*server)(nil)

type server struct {
	router chi.Router
	db     *bun.DB
}

func New(db *bun.DB) *server {
	r := chi.NewRouter()

	srv := &server{
		router: r,
		db:     db,
	}

	srv.registerMiddleware()

	srv.registerHandlers()

	return srv
}

// Public Methods
// ------------------------------------------------------------------------

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

// Private Functions and Procedures
// ------------------------------------------------------------------------
