package interceptors

import (
	"net/http"

	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
)

// TracingInterceptor inject SpanContext
type TracingInterceptor struct {
	Propagator  propagation.HTTPFormat
	Next        http.Handler
	ServiceName string
}

func (i *TracingInterceptor) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	spanCtx, ok := i.Propagator.SpanContextFromRequest(req)
	if !ok {
		i.Next.ServeHTTP(resp, req)
		return
	}

	ctx, span := trace.StartSpanWithRemoteParent(req.Context(), i.ServiceName, spanCtx)
	req = req.WithContext(ctx)
	defer span.End()

	i.Next.ServeHTTP(resp, req)
}
