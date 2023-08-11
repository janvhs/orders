package server

import (
	"github.com/go-chi/chi/middleware"
)

func (s *server) registerMiddleware() {
	// Give each request a unique ID
	s.router.Use(middleware.RequestID)
	// Get the client's ip even when proxied
	s.router.Use(middleware.RealIP)
	// Remove multiple slashes from the requested resource path
	s.router.Use(middleware.CleanPath)
	// Remove any trailing slash from the requested resource path
	s.router.Use(middleware.StripSlashes)

	// TODO: Add some kind of rate limiting (uber)
	// TODO: Add CORS (rs/cors)
	// TODO: Add CSRF protection
	// TODO: Add security headers (rs/secure)
	// TODO: Add compression
	// TODO: Add double slash removal
	// TODO: Add trailing slash removal
	// TODO: Add request timeout from

	// TODO: use charmbracelet/log middleware
	// Log every incoming request
	s.router.Use(middleware.Logger)
	// A panic should not quit the program
	s.router.Use(middleware.Recoverer)
}
