package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func New(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := log.With(
			slog.String("component", "loggingMiddleware"),
		)
		log.Info("logging middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("URL", r.URL.Path),
				slog.String("userAddress", r.RemoteAddr),
				slog.String("userAgent", r.UserAgent()),
				slog.String("reqId", middleware.GetReqID(r.Context())),
			)

			responseWriter := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			startTime := time.Now()

			defer func() {
				entry.Info(
					"request completed",
					slog.Int("status", responseWriter.Status()),
					slog.Int("bytesWritten", responseWriter.BytesWritten()),
					slog.String("duration", time.Since(startTime).String()),
				)
			}()

			next.ServeHTTP(responseWriter, r)
		}

		return http.HandlerFunc(fn)
	}
}
