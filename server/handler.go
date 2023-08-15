package server

import (
	"embed"
	"io/fs"
	"net/http"

	"git.bode.fun/orders/order"
)

func (s *server) registerHandlers() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		s.renderer.HTML(w, http.StatusOK, "index", nil)
	})
	s.router.Route("/order", order.RegisterHandlers(s.db, s.renderer))

	s.router.Get("/static/*", mustRegisterStaticHandler(s.staticFS))
}

// TODO: Add development mode, where the templates are not embedded
func mustRegisterStaticHandler(staticFS embed.FS) http.HandlerFunc {
	staticSubFS, err := fs.Sub(staticFS, "web")
	if err != nil {
		panic(err)
	}

	httpFS := http.FS(staticSubFS)
	staticHandler := http.StripPrefix("/", http.FileServer(httpFS))
	return staticHandler.ServeHTTP
}
