package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	RequestLatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_latency_histogram",
			Help:    "Histogram of HTTP request latencies in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	TasksCount = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "tasks_count",
			Help: "Current number of tasks in the system",
		},
	)
)