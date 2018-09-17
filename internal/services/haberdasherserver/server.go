package haberdasherserver

import (
	"context"
	"log"
	"math/rand"

	"github.com/nvcnvn/twirp_example/rpc/haberdasher"
	"github.com/nvcnvn/twirp_example/rpc/supervisor"
	"github.com/twitchtv/twirp"
)

// Server implements the Haberdasher service
type Server struct {
	SupervisorClient supervisor.Supervisor
}

// MakeHat implementation
func (s *Server) MakeHat(ctx context.Context, size *haberdasher.Size) (hat *haberdasher.Hat, err error) {
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}

	resp, err := s.SupervisorClient.Track(ctx, &supervisor.TrackRequest{
		Total: 1,
		Type:  supervisor.TrackRequest_HAT,
	})

	if err != nil {
		log.Println("error when calling supervisor.Track", err)
	} else {
		log.Println("supervisor acked with response", resp)
	}

	return &haberdasher.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}
