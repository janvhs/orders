package server

import (
	"embed"
	"io/fs"
	"net/http"

	"git.bode.fun/orders/order"
)

// Register routes here
// ------------------------------------------------------------------------

func (s *server) registerHandlers() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		s.renderer.HTML(w, http.StatusOK, "index", nil)
	})

	s.router.Route("/order", order.RegisterHandlers(s.db, s.renderer))

	s.router.Get("/static/*", mustRegisterStaticHandler(s.staticFS, s.isDevelopment))
}

// Private Functions and Procedures
// ------------------------------------------------------------------------

// TODO: Add development mode, where the templates are not embedded
// TODO: I HATE how passing around the fs is required. I want to get rid of that
func mustRegisterStaticHandler(staticFS embed.FS, isDevelopment bool) http.HandlerFunc {
	if isDevelopment {
		return http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))).ServeHTTP
	}

	staticSubFS, err := fs.Sub(staticFS, "web")
	if err != nil {
		panic(err)
	}

	httpFS := http.FS(staticSubFS)
	staticHandler := http.StripPrefix("/", http.FileServer(httpFS))
	return staticHandler.ServeHTTP
}
