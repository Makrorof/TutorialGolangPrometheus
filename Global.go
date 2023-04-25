package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

func init() {
	log.Println("Init")
	prometheus.Register(totalRequests)
}

func setMetricsHandle() {
	http.Handle("/metrics", promhttp.Handler())
}
