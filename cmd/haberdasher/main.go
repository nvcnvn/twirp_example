package main

import (
	"net/http"
	"time"

	"github.com/twitchtv/twirp"

	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.opencensus.io/zpages"

	"github.com/nvcnvn/twirp_example/internal/hooks"
	"github.com/nvcnvn/twirp_example/internal/services/haberdasherserver"
	"github.com/nvcnvn/twirp_example/rpc/haberdasher"
	"github.com/nvcnvn/twirp_example/rpc/supervisor"
	"go.opencensus.io/plugin/ochttp"
)

func main() {
	supervisorClient := supervisor.NewSupervisorProtobufClient("http://localhost:9991", &http.Client{Transport: &ochttp.Transport{}})

	haberdasherHandler := haberdasher.NewHaberdasherServer(&haberdasherserver.Server{
		SupervisorClient: supervisorClient,
	}, &twirp.ServerHooks{
		RequestRouted: hooks.StartSpanWhenRequestRouted,
		ResponseSent:  hooks.EndSpanWhenResponseSent,
		Error:         hooks.SetSpanStatusOnError,
	})

	mux := http.NewServeMux()
	mux.Handle(haberdasher.HaberdasherPathPrefix, haberdasherHandler)
	zpages.Handle(mux, "/debug")

	// Register stats and trace exporters to export the collected data.
	exporter := &exporter.PrintExporter{}
	trace.RegisterExporter(exporter)
	view.RegisterExporter(exporter)

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	view.SetReportingPeriod(1 * time.Second)

	http.ListenAndServe(":9990", mux)
}
