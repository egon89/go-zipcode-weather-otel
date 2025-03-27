package util

import (
	"context"
	"runtime"

	"github.com/egon89/go-zipcode-weather-gateway/internal/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var Tracer = otel.GetTracerProvider().Tracer("")

func StartSpan(ctx context.Context) (context.Context, trace.Span) {
	pc, _, _, _ := runtime.Caller(1)
	details := runtime.FuncForPC(pc)

	return Tracer.Start(ctx, details.Name())
}

func RequestIdToAttribute(ctx context.Context) attribute.KeyValue {
	return attribute.String("request-id", ctx.Value(middleware.CtxKey(0)).(string))
}
