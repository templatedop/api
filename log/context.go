package log

import (
	"context"
	"github.com/rs/zerolog"
	
)

func CtxLogger(ctx context.Context) *Logger {
	fields := make(map[string]interface{})
	
	if len(fields) > 0 {
		logger := zerolog.Ctx(ctx).With().Fields(fields).Logger()

		return &Logger{&logger}
	}

	return &Logger{zerolog.Ctx(ctx)}
}
