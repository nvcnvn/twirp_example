package main

import (
	"net/http"

	"github.com/nvcnvn/twirp_example/internal/services/haberdasherserver"
	"github.com/nvcnvn/twirp_example/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server, nil)

	http.ListenAndServe(":9990", twirpHandler)
}
