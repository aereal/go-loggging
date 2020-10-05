package logging

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type tCtxKey int

var ctxKey = tCtxKey(0)

func FromContext(ctx context.Context) (context.Context, *zap.Logger) {
	if logger, ok := ctx.Value(ctxKey).(*zap.Logger); ok {
		return ctx, logger
	}
	logger, err := newLogger()
	if err != nil {
		return ctx, nil
	}
	return withLogger(ctx, logger), logger
}

func withLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKey, logger)
}

func newLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build(zap.WithCaller(true))
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func WithEmptyLogger(ctx context.Context) context.Context {
	return withLogger(ctx, zap.NewNop())
}
