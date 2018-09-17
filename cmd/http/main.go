package main

import (
	"net/http"

	"github.com/nvcnvn/twirp_example/internal/services/haberdasherserver"
	"github.com/nvcnvn/twirp_example/rpc/haberdasher"

	"github.com/nvcnvn/twirp_example/internal/services/supervisorserver"
	"github.com/nvcnvn/twirp_example/rpc/supervisor"
)

func main() {
	supervisorHandler := supervisor.NewSupervisorServer(supervisorserver.NewServer(), nil)

	supervisorClient := supervisor.NewSupervisorProtobufClient("http://localhost:9990", &http.Client{})
	haberdasherHandler := haberdasher.NewHaberdasherServer(&haberdasherserver.Server{
		SupervisorClient: supervisorClient,
	}, nil)

	mux := http.NewServeMux()
	mux.Handle(haberdasher.HaberdasherPathPrefix, haberdasherHandler)
	mux.Handle(supervisor.SupervisorPathPrefix, supervisorHandler)

	http.ListenAndServe(":9990", mux)
}
