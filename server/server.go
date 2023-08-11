package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

var _ http.Handler = (*server)(nil)

type server struct {
	router chi.Router
	db     *sqlx.DB
}

func New(db *sqlx.DB) *server {
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
