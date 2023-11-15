package logger

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type OtelZapLogger struct {
	sugaredLogger *zap.SugaredLogger
	level         zapcore.Level
	enableTrace   bool
}

func newOtelZapLogger(config Config) Logger {
	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	writer := zapcore.Lock(os.Stdout)

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(config.JSONFormat), writer, level),
	)

	logger := zap.New(core,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()

	return &OtelZapLogger{
		sugaredLogger: logger,
		level:         level,
		enableTrace:   config.EnableTrace,
	}
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time" // This will change the key from ts to time
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (l *OtelZapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *OtelZapLogger) DebugContext(ctx context.Context, args ...interface{}) {
	l.sugaredLogger.Debug(args...)
	l.log(ctx, zap.DebugLevel, "", args)
}

func (l *OtelZapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *OtelZapLogger) DebugfContext(ctx context.Context, format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
	l.log(ctx, zap.DebugLevel, format, args)
}

func (l *OtelZapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *OtelZapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *OtelZapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *OtelZapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *OtelZapLogger) InfoContext(ctx context.Context, args ...interface{}) {
	l.sugaredLogger.Info(args...)
	l.log(ctx, zap.InfoLevel, "", args)
}

func (l *OtelZapLogger) InfofContext(ctx context.Context, format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
	l.log(ctx, zap.InfoLevel, format, args)
}

func (l *OtelZapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *OtelZapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *OtelZapLogger) WarnContext(ctx context.Context, args ...interface{}) {
	l.sugaredLogger.Warn(args...)
	l.log(ctx, zap.WarnLevel, "", args)
}

func (l *OtelZapLogger) WarnfContext(ctx context.Context, format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
	l.log(ctx, zap.WarnLevel, format, args)
}

func (l *OtelZapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *OtelZapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *OtelZapLogger) ErrorContext(ctx context.Context, args ...interface{}) {
	l.sugaredLogger.Error(args...)
	l.log(ctx, zap.ErrorLevel, "", args)
}

func (l *OtelZapLogger) ErrorfContext(ctx context.Context, format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
	l.log(ctx, zap.ErrorLevel, format, args)
}

func (l *OtelZapLogger) log(ctx context.Context, lvl zapcore.Level, template string, fmtArgs []interface{}) {
	if !l.enableTrace {
		return
	}
	// If logging at this level is completely disabled, skip the overhead of
	// string formatting.
	if lvl < zapcore.DPanicLevel && lvl < l.level {
		return
	}

	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}
	msg := getMessage(template, fmtArgs)

	if lvl >= zap.ErrorLevel {
		span.SetStatus(codes.Error, msg)
	} else {
		span.AddEvent(lvl.CapitalString() + " " + msg)
	}
}

// getMessage format with Sprint, Sprintf, or neither.
func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}
