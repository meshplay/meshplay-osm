package tracing

import (
	"context"

	"go.opentelemetry.io/otel/api/kv"
	apitrace "go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// KeyValue is the key value datastructure for store
type KeyValue struct {
	Key   string
	Value string
}

// Handler is the handler interface for tracing
type Handler interface {
	Tracer(name string) interface{}
	Span(ctx context.Context)
	AddEvent(name string, attrs ...*KeyValue)
}

// handler is the handler object for tracing
type handler struct {
	provider apitrace.Provider
	context  context.Context
	span     apitrace.Span
}

// New initiates the tracing provider for a given service
func New(service string, endpoint string) (Handler, error) {
	provider, flush, err := jaeger.NewExportPipeline(
		jaeger.WithCollectorEndpoint(endpoint),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: service,
			Tags: []kv.KeyValue{
				kv.Key("name").String(service),
				kv.Key("exporter").String("jaeger"),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		flush()
		return nil, err
	}

	return &handler{
		provider: provider,
	}, nil
}

// Tracer initiates a tracer instance
func (h *handler) Tracer(name string) interface{} {
	return h.provider.Tracer(name)
}

// Span creates a new span
func (h *handler) Span(ctx context.Context) {
	h.span = apitrace.SpanFromContext(ctx)
	h.context = ctx
}

// AddEvent adds a new event to the span
func (h *handler) AddEvent(name string, attrs ...*KeyValue) {
	kvstore := make([]kv.KeyValue, 0)
	for _, attr := range attrs {
		kvstore = append(kvstore, kv.String(attr.Key, attr.Value))
	}

	h.span.AddEvent(h.context, name, kvstore...)
}
