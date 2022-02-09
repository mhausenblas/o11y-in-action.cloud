package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// traceProvider returns Otel configured for Jaeger.
// based on:
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	environment := "test"
	service := "echosvc"
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
		)),
	)
	return tp, nil
}

func handleEcho(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	log.WithFields(log.Fields{
		"service": "echo",
	}).Info("Got input ", message)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tr := otel.Tracer("http-api")
	_, span := tr.Start(ctx, "echo")
	defer span.End()
	if rand.Intn(100) > 90 {
		log.WithFields(log.Fields{
			"service": "echo",
		}).Error("Something really bad happened :(")
		invokes.WithLabelValues(strconv.Itoa(http.StatusInternalServerError)).Inc()
		span.SetAttributes(attribute.Key("http-status-code").String(strconv.Itoa(http.StatusInternalServerError)))
		w.WriteHeader(500)
		return
	}
	invokes.WithLabelValues(strconv.Itoa(http.StatusOK)).Inc()
	span.SetAttributes(attribute.Key("http-status-code").String(strconv.Itoa(http.StatusOK)))
	fmt.Fprintf(w, message)
}

var (
	registry = prometheus.NewRegistry()
	invokes  = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "echo_total",
			Help: "Total invocations of the echo service endpoint.",
		},
		[]string{"http_status_code"},
	)
	tp = tracesdk.TracerProvider{}
)

func main() {
	rand.Seed(time.Now().UnixNano())
	registry.MustRegister(invokes)
	log.SetFormatter(&log.JSONFormatter{})
	http.HandleFunc("/echo", handleEcho)
	http.Handle("/metrics", promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	log.Fatal(http.ListenAndServe(":4242", nil))
}
