package haberdasherserver

import (
	"context"
	"math/rand"

	"github.com/nvcnvn/twirp_example/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

// Server implements the Haberdasher service
type Server struct{}

// MakeHat implementation
func (s *Server) MakeHat(ctx context.Context, size *haberdasher.Size) (hat *haberdasher.Hat, err error) {
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}
	return &haberdasher.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}
