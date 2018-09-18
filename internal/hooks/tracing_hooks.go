package hooks

import (
	"context"

	"errors"

	"github.com/twitchtv/twirp"
	"go.opencensus.io/trace"
)

var errNoMethodName = errors.New("no twirp method name found in context")

// StartSpanWhenRequestRouted must be used with twirp ServerHooks.RequestRouted
func StartSpanWhenRequestRouted(ctx context.Context) (context.Context, error) {
	methodName, ok := twirp.MethodName(ctx)
	if !ok {
		return ctx, errNoMethodName
	}

	ctx, _ = trace.StartSpan(ctx, methodName)
	return ctx, nil
}

// EndSpanWhenResponseSent must be used with twirp ServerHooks.ResponseSent
func EndSpanWhenResponseSent(ctx context.Context) {
	span := trace.FromContext(ctx)
	if span == nil {
		return
	}

	span.End()
}

// SetSpanStatusOnError  must be used with twirp ServerHooks.Error
func SetSpanStatusOnError(ctx context.Context, err twirp.Error) context.Context {
	span := trace.FromContext(ctx)
	if span == nil {
		return ctx
	}

	span.SetStatus(trace.Status{
		Code:    int32(twirp.ServerHTTPStatusFromErrorCode(err.Code())),
		Message: err.Msg(),
	})

	return ctx
}
