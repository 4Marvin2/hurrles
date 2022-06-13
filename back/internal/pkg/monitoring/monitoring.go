package monitoring

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PromMetrics struct {
	Hits    *prometheus.CounterVec
	Timings *prometheus.HistogramVec
}

func RegisterMetrics(r *mux.Router, instanseName string) *PromMetrics {
	var metrics PromMetrics

	metrics.Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_hits", instanseName),
	}, []string{"status", "path", "method"})

	metrics.Timings = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: fmt.Sprintf("%s_timings", instanseName),
		},
		[]string{"status", "path", "method"},
	)

	prometheus.MustRegister(metrics.Hits, metrics.Timings)

	r.Handle("/metrics", promhttp.Handler())

	return &metrics
}

func Metrics(metrics *PromMetrics) (mw func(http.Handler) http.Handler) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			if r.URL.Path != "/metrics" {
				metrics.Hits.WithLabelValues("requests", r.Method, "api").Inc()
			}
		})
	}
}
