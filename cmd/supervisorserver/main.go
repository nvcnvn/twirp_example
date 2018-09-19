package main

import (
	"net/http"
	"time"

	"github.com/twitchtv/twirp"
	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.opencensus.io/zpages"

	"github.com/nvcnvn/twirp_example/internal/hooks"
	"github.com/nvcnvn/twirp_example/internal/interceptors"
	"github.com/nvcnvn/twirp_example/internal/services/supervisorserver"
	"github.com/nvcnvn/twirp_example/rpc/supervisor"
)

func main() {
	supervisorServer := supervisorserver.NewServer()
	supervisorHandler := supervisor.NewSupervisorServer(supervisorServer, &twirp.ServerHooks{
		RequestRouted: hooks.StartSpanWhenRequestRouted,
		ResponseSent:  hooks.EndSpanWhenResponseSent,
		Error:         hooks.SetSpanStatusOnError,
	})

	mux := http.NewServeMux()
	mux.Handle(supervisor.SupervisorPathPrefix, &interceptors.TracingInterceptor{
		ServiceName: supervisor.SupervisorPathPrefix,
		Next:        supervisorHandler,
		Propagator:  &b3.HTTPFormat{},
	})
	zpages.Handle(mux, "/debug")

	// Register stats and trace exporters to export the collected data.
	exporter := &exporter.PrintExporter{}
	trace.RegisterExporter(exporter)
	view.RegisterExporter(exporter)

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	view.SetReportingPeriod(1 * time.Second)

	http.ListenAndServe(":9991", mux)
}
