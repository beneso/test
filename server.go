package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const addr string = ":8080"

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func recordMetrics() {
	go func() {
		for {
			request_count.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	request_count = promauto.NewCounter(prometheus.CounterOpts{
		Name: "request_count",
		Help: "The cumulative number of requests",
	})
)

func main() {
	// prometheus
	recordMetrics()

	// Expose the metrics on /metrics on port 8080.
	http.Handle("/metrics", promhttp.Handler())

	// /health
	http.HandleFunc("/health", healthHandler)

	// main content
	fs := http.FileServer(http.Dir("./content"))
	http.Handle("/", fs)
	request_count.Inc()

	// server
	log.Printf("http server starting at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}

}
