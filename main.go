package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type formatKey struct{}

func formatDecider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		format, exists := ctx.Value(formatKey{}).(string)

		// If the format already exists, bail out
		if exists {
			log.Println("format already exists:", format)
			next.ServeHTTP(w, r)
			return
		}

		// If the format does not exist, set it to text/html by default
		log.Println("setting format as:", "text/html")
		ctx = context.WithValue(ctx, formatKey{}, "text/html")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(formatDecider)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		format, exists := ctx.Value("format").(string)
		// Set the format to text/html by default
		if !exists {
			format = "text/html"
		}

		log.Println("Format:", format)
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8080", r)
}
