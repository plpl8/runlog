package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "runlog_http_requests_total",
			Help: "Total number of HTTP requests handled by RunLog.",
		},
		[]string{"method", "route", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "runlog_http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)

	runsCreatedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "runlog_runs_created_total",
			Help: "Total number of successfully created runs.",
		},
	)

	runsUpdatedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "runlog_runs_updated_total",
			Help: "Total number of successfully updated runs.",
		},
	)

	runsDeletedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "runlog_runs_deleted_total",
			Help: "Total number of successfully deleted runs.",
		},
	)

	runsDistanceTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "runlog_runs_distance_kilometers_total",
			Help: "Cumulative distance of successfully created runs, in kilometers.",
		},
	)
)

func init() {
	prometheus.MustRegister(
		httpRequestsTotal,
		httpRequestDuration,
		runsCreatedTotal,
		runsUpdatedTotal,
		runsDeletedTotal,
		runsDistanceTotal,
	)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		route := c.FullPath()
		if route == "" {
			route = "unmatched"
		}

		httpRequestsTotal.WithLabelValues(
			c.Request.Method,
			route,
			strconv.Itoa(c.Writer.Status()),
		).Inc()

		httpRequestDuration.WithLabelValues(
			c.Request.Method,
			route,
		).Observe(time.Since(start).Seconds())
	}
}

func MetricsHandler() gin.HandlerFunc {
	handler := promhttp.Handler()

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
