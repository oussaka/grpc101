// profile/profile.go
package profile

import (
	"log"

	"github.com/ensep/grpc101/proto/profile"
	"golang.org/x/net/context"
)

type Server struct {
	profile.UnimplementedProfileServer
}

func (s *Server) Create(ctx context.Context, req *profile.CreateRequest) (*profile.CreateResponse, error) {
	log.Printf("Receive message body from client: %s", req.Name)
	return &profile.CreateResponse{Message: "Profile Created!"}, nil
}
