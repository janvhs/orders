package server

import (
	"net/http"

	"git.bode.fun/orders/order"
)

func (s *server) registerHandlers() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		s.renderer.HTML(w, http.StatusOK, "index", nil)
	})
	s.router.Route("/order", order.RegisterHandlers(s.db, s.renderer))

}

// TODO: Static files but this is not working yet
// func setUpStaticFileServer(r chi.Router, staticFS embed.FS, isDevelopment bool) {
// 	var handler http.Handler

// 	if isDevelopment {
// 		handler = http.FileServer(http.Dir("web/static"))
// 	} else {
// 		handler = http.FileServer(http.FS(staticFS))
// 	}

// 	r.Get("/static/*", http.StripPrefix("web/static", handler).ServeHTTP)
// }
