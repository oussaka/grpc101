package handler

import (
	"log"

	"github.com/ensep/grpc101/proto/profile"
	"golang.org/x/net/context"
)

type Profile struct {
	profile.UnimplementedProfileServer
}

func (s *Profile) Create(ctx context.Context, req *profile.CreateRequest) (*profile.CreateResponse, error) {
	log.Printf("Receive message body from client: %s", req.GetName())
	return &profile.CreateResponse{Message: "Profile Created!"}, nil
}
