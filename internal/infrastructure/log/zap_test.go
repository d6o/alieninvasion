package log_test

import (
	"github.com/d6o/alieninvasion/internal/infrastructure/log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewZap(t *testing.T) {
	testCases := []struct {
		name    string
		verbose bool
		level   zapcore.Level
		err     error
	}{
		{
			name:    "Test new zap with verbose false",
			verbose: false,
			level:   zapcore.FatalLevel,
			err:     nil,
		},
		{
			name:    "Test new zap with verbose true",
			verbose: true,
			level:   zapcore.DebugLevel,
			err:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger, err := log.NewZap(tc.verbose)
			assert.Equal(t, tc.err, err)

			if err == nil {
				assert.NotNil(t, logger)
				assert.True(t, logger.Core().Enabled(tc.level))
			}
		})
	}
}

func TestNewZapNop(t *testing.T) {
	logger := log.NewZapNop()
	assert.NotNil(t, logger)
	assert.Equal(t, zap.NewNop(), logger)
}
