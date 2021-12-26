package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type formatKey struct{}

func formatDecider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		format, exists := ctx.Value(formatKey{}).(string)

		// If the format already exists, bail out
		if exists {
			log.Trace().Msgf("format already exists: %s", format)
			next.ServeHTTP(w, r)
			return
		}

		// If the format does not exist, set it to text/html by default
		log.Trace().Msgf("setting format as: %s", "text/html")
		ctx = context.WithValue(ctx, formatKey{}, "text/html")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(formatDecider)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		format := ctx.Value(formatKey{}).(string)

		log.Print("Format:", format)

		w.Write([]byte(fmt.Sprintln("PONG")))
		w.Write([]byte(fmt.Sprintln("format:", format)))
	})

	addPostfixRoutes(r, map[string]string{
		".json": "application/json",
		".html": "text/html",
	})
	logRoutes(r)

	http.ListenAndServe(":8080", r)
}

// Add extra routes to override the return format via path file extensions
func addPostfixRoutes(r *chi.Mux, formats map[string]string) {
	chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		for extension, format := range formats {
			r.
				// Inline middleware to overwrite the format
				With(overwriteFormatHandler(format)).
				// Register the original handler to the route path plus the new extension
				Method(method, route+extension, handler)
		}

		return nil
	})
}

// Middleware to overwrite the format value of a Context
func overwriteFormatHandler(format string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log.Trace().Msgf("overridding format via path extension: %s", format)
			ctx = context.WithValue(ctx, formatKey{}, format)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Log out routes
func logRoutes(r *chi.Mux) {
	log.Warn().Msg("Registered routes:")
	chi.Walk(r, func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		log.Warn().Msgf(" %s %s", method, route)
		return nil
	})
}
