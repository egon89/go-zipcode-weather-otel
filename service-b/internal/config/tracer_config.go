package config

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func InitTracer() func() {
	log.Printf("initializing tracing to %s\n", OtelCollectorHost)
	ctx := context.Background()

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(OtelCollectorHost),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("error creating otlp trace exporter in %s: %v", OtelCollectorHost, err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			attribute.String("service.name", "service-b"),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("error finishing tracing: %v", err)
		}
	}
}
