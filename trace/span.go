package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Span struct {
	SpanName string
	span     trace.Span
}

func (s *Span) Event(ctx context.Context) {
	tracerProvider := otel.GetTracerProvider()
	_, span := tracerProvider.Tracer(s.SpanName).Start(ctx, s.SpanName)
	s.span = span
}

func (s Span) Close() {
	s.span.End()
}
