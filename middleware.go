package main

import (
	"context"
	"net/http"
	"time"

	"github.com/elnormous/contenttype"
	"github.com/go-chi/chi/v5/middleware"
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

		// Parse the accept header to detect format
		// Defaults to first media type if none is provided
		availableMediaTypes := []contenttype.MediaType{
			contenttype.NewMediaType("text/html"),
			contenttype.NewMediaType("application/json"),
		}
		accepted, _, _ := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
		log.Trace().Msgf("setting format as: %s", accepted.String())
		ctx = context.WithValue(ctx, formatKey{}, accepted.String())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func requestLogger() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			h.ServeHTTP(ww, r)

			difference := time.Since(t1)

			log.Info().
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Int("status", ww.Status()).
				Int("bytesWritten", ww.BytesWritten()).
				Dur("duration", difference).
				Msgf("%s %s - %d %s", r.Method, r.URL.Path, ww.Status(), difference)
		})
	}
}
