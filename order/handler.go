package order

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
	"github.com/uptrace/bun"
)

func RegisterHandlers(db *bun.DB, render *render.Render) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", handleIndexOrder(db, render))
	}
}

func handleIndexOrder(db *bun.DB, render *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create request specific Repository and Service
		repo := NewRepo(db)
		serv := NewService(repo)
		orders, err := serv.ListAll()

		if err != nil {
			render.Text(w, http.StatusInternalServerError, err.Error())
			return
		}

		render.HTML(w, http.StatusOK, "order/index", orders)

	}
}
