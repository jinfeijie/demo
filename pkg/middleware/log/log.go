package log

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Logger log interface
type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
}

var logger Logger = zap.NewExample()

func GetLog() Logger {
	return logger
}

func L() Logger {
	return GetLog()
}

func SetLog(l Logger) {
	logger = l
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = getTraceId(ctx, fields...)
	GetLog().Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = getTraceId(ctx, fields...)
	GetLog().Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = getTraceId(ctx, fields...)
	GetLog().Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = getTraceId(ctx, fields...)
	GetLog().Error(msg, fields...)
}

func DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	fields = getTraceId(ctx, fields...)
	GetLog().DPanic(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = getTraceId(ctx, fields...)
	GetLog().Fatal(msg, fields...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	fields = getTraceId(ctx, fields...)
	GetLog().Panic(msg, fields...)
}

func getTraceId(ctx context.Context, fields ...zap.Field) []zap.Field {
	traceId := ctx.Value("__traceId__")
	if traceId == nil {
		traceId = uuid.New().String()
		ctx = context.WithValue(ctx, "__traceId__", traceId)
	}
	traceIdStr := traceId.(string)
	fields = append(fields, zap.String("traceId", traceIdStr))
	return fields
}
