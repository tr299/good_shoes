package logger

import (
	"context"
)

// Logger
type Logger interface {
	Debug(args ...interface{})

	DebugContext(ctx context.Context, args ...interface{})

	Debugf(format string, args ...interface{})

	DebugfContext(ctx context.Context, format string, args ...interface{})

	Info(args ...interface{})

	Infof(format string, args ...interface{})

	InfoContext(ctx context.Context, args ...interface{})

	InfofContext(ctx context.Context, format string, args ...interface{})

	Warn(args ...interface{})

	Warnf(format string, args ...interface{})

	WarnContext(ctx context.Context, args ...interface{})

	WarnfContext(ctx context.Context, format string, args ...interface{})

	Error(args ...interface{})

	Errorf(format string, args ...interface{})

	ErrorContext(ctx context.Context, args ...interface{})

	ErrorfContext(ctx context.Context, format string, args ...interface{})

	Fatal(args ...interface{})

	Fatalf(format string, args ...interface{})
}

type ErrorLog struct {
	data    any
	message string
}

// Config
type Config struct {
	Level       string
	JSONFormat  bool
	EnableTrace bool
}

var log Logger

// init
// auto init,this can useful to use zap/log directly
func init() {
	log = newOtelZapLogger(Config{
		Level:       "debug",
		EnableTrace: true,
	})
}

func Init(config Config) {
	log = newOtelZapLogger(config)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func DebugContext(ctx context.Context, args ...interface{}) {
	log.DebugContext(ctx, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	log.DebugfContext(ctx, format, args...)
}

func DebugfContext(ctx context.Context, format string, args ...interface{}) {
	log.DebugfContext(ctx, format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func InfoContext(ctx context.Context, args ...interface{}) {
	log.InfoContext(ctx, args...)
}

func InfofContext(ctx context.Context, format string, args ...interface{}) {
	log.InfofContext(ctx, format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func WarnContext(ctx context.Context, args ...interface{}) {
	log.WarnContext(ctx, args...)
}

func WarnfContext(ctx context.Context, format string, args ...interface{}) {
	log.WarnfContext(ctx, format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func ErrorContext(ctx context.Context, args ...interface{}) {
	log.ErrorContext(ctx, args...)
}

func ErrorfContext(ctx context.Context, format string, args ...interface{}) {
	log.ErrorfContext(ctx, format, args...)
}
