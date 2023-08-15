package server

import (
	"github.com/go-chi/chi/v5/middleware"
)

// Add middleware here
// ------------------------------------------------------------------------

func (s *server) registerMiddleware() {
	// Give each request a unique ID
	s.router.Use(middleware.RequestID)
	// Get the client's ip even when proxied
	s.router.Use(middleware.RealIP)
	// Remove multiple slashes from the requested resource path
	s.router.Use(middleware.CleanPath)
	// Remove any trailing slash from the requested resource path
	s.router.Use(middleware.StripSlashes)

	// Paths are clean and ready to be logged :)

	// Log every incoming request
	// Log middleware depends on Recover
	s.router.Use(middleware.Logger)
	// A panic should not quit the program
	s.router.Use(middleware.Recoverer)

	// FIXME: When needed, add the following CORS middleware: "github.com/rs/cors"
	// FIXME: Add security headers via "github.com/unrolled/secure"
	// FIXME: Add uber rate-limit middleware
	// FIXME: Add CSRF protection via "github.com/justinas/nosurf"
	// FIXME: Add Gzip via github.com/klauspost/compress
}
