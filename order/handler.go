package order

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func RegisterHandlers(db *sqlx.DB) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", handleIndexOrder(db))
	}
}

func handleIndexOrder(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create request specific Repository and Service
		repo := NewRepo(db)
		serv := NewService(repo)
		orders, err := serv.ListAll()

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, err)
			return
		}

		w.WriteHeader(200)
		fmt.Fprintln(w, orders)
	}
}
