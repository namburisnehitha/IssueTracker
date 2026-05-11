package telemetry

import "go.opentelemetry.io/otel/metric"

type Metrics struct {
	Counter   metric.Int64Counter
	Histogram metric.Float64Histogram
}

func NewMetrics(meter metric.Meter) (Metrics, error) {
	counter, err := meter.Int64Counter("http.server.request.count")
	if err != nil {
		return Metrics{}, err
	}
	histogram, err := meter.Float64Histogram("http.server.request.duration")
	if err != nil {
		return Metrics{}, err
	}
	return Metrics{
		Counter:   counter,
		Histogram: histogram,
	}, nil
}
