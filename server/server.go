package server

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
	"github.com/uptrace/bun"
)

var _ http.Handler = (*server)(nil)

// Implementation
// ------------------------------------------------------------------------

type server struct {
	isDevelopment bool
	router        chi.Router
	renderer      *render.Render
	db            *bun.DB
	staticFS      embed.FS
}

func New(db *bun.DB, isDevelopment bool, templateFS, staticFS embed.FS) *server {
	r := chi.NewRouter()

	var renderFS render.FileSystem = &render.EmbedFileSystem{
		FS: templateFS,
	}

	if isDevelopment {
		renderFS = &render.LocalFileSystem{}
	}

	// TODO: Remove render, because it seems to be an unnecessary dependency
	renderer := render.New(render.Options{
		IsDevelopment: isDevelopment,
		Layout:        "layouts/index",
		Directory:     "web/templates",
		Extensions:    []string{".html"}, // "go.html" or "tmpl.html", sadly do not work
		FileSystem:    renderFS,
		// FIXME: If https://github.com/Masterminds/sprig is needed, add it here
		// Funcs: sprig.FuncMap(),
	})

	srv := &server{
		router:        r,
		db:            db,
		isDevelopment: isDevelopment,
		renderer:      renderer,
		staticFS:      staticFS,
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
