package telemetry

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitTraces() {
	exp, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithEndpoint("datadog-agent.datadog.svc.cluster.local:4317"),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
	)
	otel.SetTracerProvider(tp)
}
