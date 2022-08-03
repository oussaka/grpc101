// server.go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ensep/grpc101/proto/profile"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC 101 Teknasyon!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := profile.Server{}

	grpcServer := grpc.NewServer()

	profile.RegisterProfileServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
