package ctxvaluer

import (
	"context"

	"github.com/templatedop/api/log"
	"github.com/templatedop/api/util/appctx"
	//"go.uber.org/zap"
)

const (
	CorrelationIDKey = "x-correlationId"
	ExecutorUserKey  = "x-executor-user"
	TraceIDKey       = "trace-id"
	SpanIDKey        = "span-id"
	AgentNameKey     = "x-agent-name"
	OwnerKey         = "x-owner"
)

var (
	CorrelationID = appctx.NewValuer[string](CorrelationIDKey)
	ExecutorUser  = appctx.NewValuer[string](ExecutorUserKey)
	TraceID       = appctx.NewValuer[string](TraceIDKey)
	SpanID        = appctx.NewValuer[string](SpanIDKey)
	AgentName     = appctx.NewValuer[string](AgentNameKey)
	Owner         = appctx.NewValuer[string](OwnerKey)
)

type CreateParams struct {
	CorrelationID string
	ExecutorUser  string
	AgentName     string
	Owner         string
}

func CreateBaseTaskContext(parent context.Context, params CreateParams, l *log.Logger) context.Context {
	ctx := parent
	ctx = CorrelationID.Set(ctx, params.CorrelationID)
	ctx = ExecutorUser.Set(ctx, params.ExecutorUser)
	ctx = AgentName.Set(ctx, params.AgentName)
	ctx = Owner.Set(ctx, params.Owner)
	//return ctx
	return log.WithLogger(ctx, l)
	// ll:=log.ToZerolog().With().
	// 	Str(CorrelationIDKey, params.CorrelationID).
	// 	Str(ExecutorUserKey, params.ExecutorUser).
	// 	Str(AgentNameKey, params.AgentName).
	// 	Str(OwnerKey, params.Owner).Logger()

	// With().
	// 	Str(CorrelationIDKey, params.CorrelationID).
	// 	Str(ExecutorUserKey, params.ExecutorUser).
	// 	Str(AgentNameKey, params.AgentName).
	// 	Str(OwnerKey, params.Owner).
	// 	logger

	// l := logger.New().With(
	// 	zap.String(CorrelationIDKey, params.CorrelationID),
	// 	zap.String(ExecutorUserKey, params.ExecutorUser),
	// 	zap.String(AgentNameKey, params.AgentName),
	// 	zap.String(OwnerKey, params.Owner),
	// )

	//return log.WithLogger(ctx, l)
	//return log.WithContext(ctx)

	//return log.WithLogger(ctx, l)
}
