package appcontext

import (
	"context"
	"github.com/d6o/alieninvasion/internal/infrastructure/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestLogger(t *testing.T) {
	emptyLogger := log.NewZapNop()

	loggerWithData, err := log.NewZap(false)
	assert.NoError(t, err)
	loggerWithData.With(zap.Int("test", 123))

	tests := []struct {
		name  string
		setup func() context.Context
		want  *zap.Logger
	}{
		{
			name: "Context is empty",
			setup: func() context.Context {
				return context.Background()
			},
			want: emptyLogger,
		},
		{
			name: "Context has a logger",
			setup: func() context.Context {
				ctx := context.Background()
				return WithLogger(ctx, loggerWithData)
			},
			want: loggerWithData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setup()
			got := Logger(ctx)
			if reflect.DeepEqual(got, tt.want) {
				return
			}
			if !assert.IsType(t, &zap.SugaredLogger{}, got) {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLogger(t *testing.T) {
	ctx := context.Background()
	logger, err := log.NewZap(false)
	assert.NoError(t, err)
	want := context.WithValue(ctx, loggerKey, logger)

	if got := WithLogger(ctx, logger); !reflect.DeepEqual(got, want) {
		t.Errorf("WithLogger() = %v, want %v", got, want)
	}
}

func Test_loggerFromContext(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() context.Context
		want    *zap.Logger
		wantErr bool
	}{
		{
			name: "Context has a logger",
			setup: func() context.Context {
				ctx := context.Background()
				logger, err := log.NewZap(false)
				assert.NoError(t, err)
				return WithLogger(ctx, logger)
			},
			want:    &zap.Logger{},
			wantErr: false,
		},
		{
			name: "Context is empty",
			setup: func() context.Context {
				return context.Background()
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setup()
			got, err := loggerFromContext(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("loggerFromContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				return
			}
			if !assert.IsType(t, &zap.Logger{}, got) {
				t.Errorf("loggerFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
