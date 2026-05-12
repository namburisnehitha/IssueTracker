package telemetry

import (
	"log/slog"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func MetricsMiddleware(logger *slog.Logger, m *Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx, span := otel.Tracer("middleware").Start(r.Context(), "http.request")
			defer span.End()
			r = r.WithContext(ctx)

			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start).Seconds()
			trace_id := trace.SpanFromContext(r.Context()).SpanContext().TraceID().String()
			logger.Info("request", "method", r.Method, "path", r.URL.Path, "duration", duration, "trace-id", trace_id)
			m.Counter.Add(r.Context(), 1)
			m.Histogram.Record(r.Context(), duration)
		})
	}
}
