package order

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/uptrace/bun"
)

func RegisterHandlers(db *bun.DB) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", handleIndexOrder(db))
	}
}

func handleIndexOrder(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create request specific Repository and Service
		repo := NewRepo(db)
		serv := NewService(repo)
		orders, err := serv.ListAll()

		jencode := json.NewEncoder(w)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintln(w, err)
			return
		}

		w.WriteHeader(200)
		jencode.Encode(orders)
	}
}
