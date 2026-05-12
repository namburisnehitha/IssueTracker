package telemetry

import (
	"net/http"
	"time"
)

func MetricsMiddleware(m *Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			m.Counter.Add(r.Context(), 1)
			m.Histogram.Record(r.Context(), time.Since(start).Seconds())
		})
	}
}
