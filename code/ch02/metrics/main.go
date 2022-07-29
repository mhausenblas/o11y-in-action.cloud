package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func handleEcho(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	log.WithFields(log.Fields{
		"service": "echo",
	}).Info("Got input ", message)
	if rand.Intn(100) > 90 {
		log.WithFields(log.Fields{
			"service": "echo",
		}).Error("Something really bad happened :(")
		invokes.WithLabelValues(strconv.Itoa(http.StatusInternalServerError)).Inc() // <5>
		w.WriteHeader(500)
		return
	}
	invokes.WithLabelValues(strconv.Itoa(http.StatusOK)).Inc() // <5>
	fmt.Fprintf(w, message)
}

var (
	registry = prometheus.NewRegistry() // <1>
	invokes  = promauto.NewCounterVec(  // <2>
		prometheus.CounterOpts{
			Name: "echo_total",
			Help: "Total invocations of the echo service endpoint.",
		},
		[]string{"http_status_code"},
	)
)

func main() {
	rand.Seed(time.Now().UnixNano())
	registry.MustRegister(invokes) // <3>
	log.SetFormatter(&log.JSONFormatter{})
	http.HandleFunc("/echo", handleEcho)
	http.Handle("/metrics", promhttp.HandlerFor( // <4>
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))
	log.Fatal(http.ListenAndServe(":4242", nil))
}
