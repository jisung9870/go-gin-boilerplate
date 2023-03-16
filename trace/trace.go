package trace

import (
	"context"
	"log"

	"github.com/JisungPark0319/go-gin-boilerplate/config"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

type Trace struct {
	Endpoint       string
	ServiceName    string
	TracerProvider *sdktrace.TracerProvider
}

func (t *Trace) Init(cfg config.TraceConfig) error {
	t.Endpoint = cfg.Endpoint
	t.ServiceName = cfg.ServiceName
	tp, err := initTracer(t.Endpoint, t.ServiceName)
	if err != nil {
		return err
	}
	t.TracerProvider = tp
	return nil
}

func (t *Trace) Provider() {
	otel.SetTracerProvider(t.TracerProvider)
}

func (t *Trace) Close() {
	if err := t.TracerProvider.Shutdown(context.Background()); err != nil {
		log.Printf("Error shutting down tracer provider: %v", err)
	}
}

func (t *Trace) GinMiddleware(engine *gin.Engine) {
	engine.Use(otelgin.Middleware(t.ServiceName))
}

func (t *Trace) GormMiddleware(db *gorm.DB) error {
	if err := db.Use(tracing.NewPlugin()); err != nil {
		return err
	}
	return nil
}

func initTracer(url string, serviceName string) (*sdktrace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)
	return tp, nil
}
