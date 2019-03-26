package logger

import (
	"context"
)

type Logger interface {
	Trace(ctx context.Context, args ...interface{})
	Tracef(ctx context.Context, format string, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
}
