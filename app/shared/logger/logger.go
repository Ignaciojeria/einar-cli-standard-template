package logger

import (
	"archetype/app/shared/configuration"
	"archetype/app/shared/constants"
	"log/slog"
	"os"
	"strconv"

	"go.opentelemetry.io/otel/trace"
)

// Datadog trace and log correlation :
// https://docs.datadoghq.com/tracing/other_telemetry/connect_logs_and_traces/opentelemetry/?tab=go
const (
	ddTraceIDKey = "dd.trace_id"
	ddSpanIDKey  = "dd.span_id"
	ddServiceKey = "dd.service"
	ddEnvKey     = "dd.env"
	ddVersionKey = "dd.version"
)

// Default opentelemetry trace and log correlation :
const (
	traceIDKey = "trace_id"
	spanIDKey  = "span_id"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func SpanLogger(span trace.Span) *slog.Logger {
	traceID := span.SpanContext().TraceID().String()
	spanID := span.SpanContext().SpanID().String()

	ddService := configuration.Values().DD_SERVICE
	ddEnv := configuration.Values().DD_ENV
	ddVersion := configuration.Values().DD_VERSION

	if ddService == "" || ddEnv == "" || ddVersion == "" {
		return logger.With(
			slog.String(traceIDKey, traceID),
			slog.String(spanIDKey, spanID),
		)
	}
	return logger.With(
		slog.String(traceIDKey, traceID),
		slog.String(spanIDKey, spanID),
		slog.String(ddTraceIDKey, convertTraceID(traceID)),
		slog.String(ddSpanIDKey, convertTraceID(spanID)),
		slog.String(ddServiceKey, ddService),
		slog.String(ddEnvKey, ddEnv),
		slog.String(ddVersionKey, ddVersion),
	)
}

type CustomLogFields map[string]interface{}

func LogSpanError(span trace.Span, message string, fields CustomLogFields) {
	SpanLogger(span).Error(message, constants.Fields, fields)
}

func LogSpanInfo(span trace.Span, message string, fields CustomLogFields) {
	SpanLogger(span).Info(message, constants.Fields, fields)
}

func Logger() *slog.Logger {
	return logger
}

func convertTraceID(id string) string {
	if len(id) < 16 {
		return ""
	}
	if len(id) > 16 {
		id = id[16:]
	}
	intValue, err := strconv.ParseUint(id, 16, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatUint(intValue, 10)
}
