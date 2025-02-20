package otel

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func InitMeterProvider() {
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal("Could not create prometheus exporter: ", err)
	}

	provider := metric.NewMeterProvider(metric.WithReader((exporter)))
	otel.SetMeterProvider(provider)
}
