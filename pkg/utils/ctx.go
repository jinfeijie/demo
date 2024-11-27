package utils

import (
	"context"
	"github.com/google/uuid"
)

func JobCtx(ctx context.Context) context.Context {
	traceId := ctx.Value("__traceId__")
	if traceId == nil {
		traceId = uuid.New().String()
		ctx = context.WithValue(ctx, "__traceId__", traceId)
	}
	return ctx
}
