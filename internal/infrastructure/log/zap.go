package log

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Config struct {
		Level string `envconfig:"LOG_LEVEL"`
	}
)

// NewZap creates and configures a Zap instance to be used by the application.
func NewZap(verbose bool) (*zap.Logger, error) {
	level := zapcore.FatalLevel
	if verbose {
		level = zapcore.DebugLevel
	}

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.Level = zap.NewAtomicLevelAt(level)

	logger, err := zapConfig.Build()
	return logger, errors.Wrap(err, "can't create zap instance, check configurations")
}

func NewZapNop() *zap.Logger {
	return zap.NewNop()
}
