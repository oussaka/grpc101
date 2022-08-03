// profile/profile.go
package profile

import (
	"log"

	"golang.org/x/net/context"
	"github.com/ensep/grpc101/proto/profile"
)

type Server struct {
}

func (s *Server) Create(ctx context.Context, req *profile.CreateRequest) (*profile.CreateResponse, error) {
	log.Printf("Receive message body from client: %s", req.Body)
	return &Message{Body: "Profile Created!"}, nil
}
