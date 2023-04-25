package main

import (
	"net/http"
	"time"
)

func main() {
	setMetricsHandle()

	go func() {
		for {
			time.Sleep(time.Second * 5)

			totalRequests.WithLabelValues("Label").Inc()
		}
	}()

	http.ListenAndServe(":9191", nil)
}
