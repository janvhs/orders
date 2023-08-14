package server

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
	"github.com/uptrace/bun"
)

var _ http.Handler = (*server)(nil)

type server struct {
	isDevelopment bool
	router        chi.Router
	renderer      *render.Render
	db            *bun.DB
}

func New(db *bun.DB, isDevelopment bool, templateFS embed.FS) *server {
	r := chi.NewRouter()

	renderer := render.New(render.Options{
		IsDevelopment: isDevelopment,
		Layout:        "layouts/index",
		Directory:     "web/templates",
		Extensions:    []string{".html"}, // "go.html" or "tmpl.html", sadly do not work
		FileSystem: &render.EmbedFileSystem{
			FS: templateFS,
		},
		UseMutexLock: true, // Im scared of data-races and do not want to create the renderer for each request
		// FIXME: If https://github.com/Masterminds/sprig is needed, add it here
		// Funcs: sprig.FuncMap(),
	})

	srv := &server{
		router:        r,
		db:            db,
		isDevelopment: isDevelopment,
		renderer:      renderer,
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
