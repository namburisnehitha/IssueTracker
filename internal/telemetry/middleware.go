package telemetry

import (
	"log/slog"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func MetricsMiddleware(logger *slog.Logger, m *Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx, span := otel.Tracer("middleware").Start(r.Context(), "http.request")
			defer span.End()
			r = r.WithContext(ctx)

			wrapped := &responseWriter{ResponseWriter: w, status: http.StatusOK}

			start := time.Now()
			next.ServeHTTP(wrapped, r)
			duration := time.Since(start).Seconds()

			span.SetAttributes(
				semconv.HTTPResponseStatusCodeKey.Int(wrapped.status),
			)

			trace_id := trace.SpanFromContext(r.Context()).SpanContext().TraceID().String()
			logger.Info("request", "method", r.Method, "path", r.URL.Path, "duration", duration, "trace-id", trace_id)
			m.Counter.Add(r.Context(), 1)
			m.Histogram.Record(r.Context(), duration)
		})
	}
}
