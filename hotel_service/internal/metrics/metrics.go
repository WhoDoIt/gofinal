package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	ResponseTime *prometheus.HistogramVec
	RequestTotal *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	metrics := Metrics{}

	buckets := []float64{0.005, 0.01, 0.025, 0.05, .1, .25, .5, 1, 2.5}
	metrics.ResponseTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "hotel",
		Name:      "http_server_response_time",
		Help:      "Measures response time in seconds",
		Buckets:   buckets,
	}, []string{"request", "method", "status_code"})

	prometheus.MustRegister(metrics.ResponseTime)

	metrics.RequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "hotel",
		Name:      "http_requests_total",
		Help:      "Measures total count of requests",
	}, []string{"request", "method", "status_code"})

	prometheus.MustRegister(metrics.RequestTotal)

	return &metrics
}
