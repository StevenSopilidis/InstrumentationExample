package otel

import (
	"context"
	"log"
	"ping/utils"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func InitTracerProvider(ctx context.Context, config utils.Config) func(context.Context) error {
	exporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(config.TracingEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatal("Could not connect to tracing endpoint: ", config.TracingEndpoint)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.ServiceName),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}
