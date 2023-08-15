// Package logger creates a logger instance
package logger

import (
	"errors"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a wrapper around `zap.Logger` that adds a few methods.
// @property  - `Logger` is the name of the struct.
type Logger struct{ *zap.Logger }

// NewLogger creates a new logger with the given level and returns it
func NewLogger(opts ...Option) (*Logger, error) {
	options := LoggerOptions{}
	for _, o := range opts {
		o(&options)
	}

	// write syncers
	stdoutSyncer := zapcore.Lock(os.Stdout)

	var l zapcore.Level

	switch options.Level {
	case "debug":
		l = zapcore.DebugLevel
	case "info":
		l = zapcore.InfoLevel
	case "error":
		l = zapcore.ErrorLevel
	case "panic":
		l = zapcore.PanicLevel
	case "fatal":
		l = zapcore.FatalLevel
	default:
		return nil, errors.New("invalid level type passed")
	}

	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(stdoutSyncer),
		zapcore.Level(l),
	)

	// finally construct the logger with the tee core
	logger := zap.New(core)
	logger = logger.Named(options.Name)
	defer logger.Sync()

	return &Logger{logger}, nil
}
