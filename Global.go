package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var serverCpuBasicSecondsTotalGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "server_cpu_basic_seconds_total",
	Help: "Server CPU basic",
}, []string{"path"})

var serverCpuUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "server_cpu_usage",
	Help: "Server CPU Usage",
})

func init() {
	log.Println("Init")

	prometheus.Register(serverCpuBasicSecondsTotalGauge)
	prometheus.Register(serverCpuUsageGauge)
}

func setMetricsHandle() {
	http.Handle("/metrics", promhttp.Handler())
}
