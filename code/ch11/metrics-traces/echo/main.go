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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// traceProvider returns OTel-configured for provider.
// See https://github.com/open-telemetry/opentelemetry-go/blob/main/example/otel-collector/main.go
func initProvider() (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("echo-svc"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "otel-col:4317",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tracerProvider.Shutdown, nil
}

func handleEcho(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	log.WithFields(log.Fields{
		"service": "echo",
	}).Info("Got input ", message)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tr := otel.Tracer("echo-api")
	_, span := tr.Start(ctx, "message")
	spanContext := span.SpanContext()
	traceID := spanContext.TraceID().String()
	defer span.End()

	if rand.Intn(100) > 60 {
		log.WithFields(log.Fields{
			"service": "echo",
		}).Error("Something really bad happened :(")
		invoke, _ := invokes.GetMetricWith(prometheus.Labels{
			"http_status_code": strconv.Itoa(http.StatusInternalServerError),
		})
		invoke.(prometheus.ExemplarAdder).AddWithExemplar(1, prometheus.Labels{
			"traceID": traceID,
		})
		span.SetAttributes(
			attribute.String("http_status_code", strconv.Itoa(http.StatusInternalServerError)),
		)
		w.WriteHeader(500)
		return
	}
	invoke, _ := invokes.GetMetricWith(prometheus.Labels{
		"http_status_code": strconv.Itoa(http.StatusOK),
	})
	invoke.(prometheus.ExemplarAdder).AddWithExemplar( // <1>
		1,
		prometheus.Labels{"traceID": traceID},
	)
	invokes.WithLabelValues( // <2>
		strconv.Itoa(http.StatusOK),
	).Inc()
	span.SetAttributes( // <3>
		attribute.String("http_status_code", strconv.Itoa(http.StatusOK)),
	)
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
)

func main() {
	shutdown, err := initProvider()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

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
	log.Fatal(http.ListenAndServe(":8888", nil))
}
