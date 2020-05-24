package log

import (
	"context"

	"go.uber.org/zap"
)

type ctxMarkerRequestId struct{}

var (
	ctxKeyRequestId = &ctxMarkerRequestId{}
)

// ExtractLogger returns the logger with a request id if exists.
func ExtractLogger(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil { // return the default logger if a context is nil
		return Logger()
	}

	requestId, ok := GetRequestId(ctx)
	if !ok || requestId == "" {
		return Logger()
	}
	return Logger().With("request_id", requestId)
}

// GetRequestId retrieves request id from the context.
func GetRequestId(ctx context.Context) (string, bool) {
	requestId, ok := ctx.Value(ctxKeyRequestId).(string)
	return requestId, ok
}

// PutRequestId puts request id into the context.
func PutRequestId(ctx context.Context, requestId string) context.Context {
	if requestId == "" {
		return ctx
	}
	return context.WithValue(ctx, ctxKeyRequestId, requestId)
}
