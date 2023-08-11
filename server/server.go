package server

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

var _ http.Handler = (*server)(nil)

type server struct {
	logger *log.Logger
	router chi.Router
	db     *sqlx.DB
}

func New(logger *log.Logger, db *sqlx.DB) *server {
	r := chi.NewRouter()

	srv := &server{
		logger: logger,
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
