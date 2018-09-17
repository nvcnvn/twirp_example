package supervisorserver

import (
	"context"

	"github.com/nvcnvn/twirp_example/rpc/supervisor"
)

// Server implements the Supervior service
type Server struct {
	Storage countMap
}

// NewServer return a Server with initialed Storage
func NewServer() *Server {
	return &Server{
		Storage: make(countMap),
	}
}

type countMap map[supervisor.TrackRequest_ProductType]int32

// Track implementation
func (s *Server) Track(ctx context.Context, req *supervisor.TrackRequest) (*supervisor.TrackRequest, error) {
	s.Storage[req.Type]++
	return &supervisor.TrackRequest{
		Total: s.Storage[req.Type],
		Type:  req.Type,
	}, nil
}
