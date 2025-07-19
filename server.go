// server.go
package main

import (
	"fmt"
	"github.com/ensep/grpc101/handler"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/reflection"
	"net"

	"github.com/ensep/grpc101/proto/profile"
	"google.golang.org/grpc"
)

const (
	GrpcServerAddress string = "0.0.0.0:9000"
)

func main() {
	fmt.Println("Go gRPC 101 Teknasyon!")

	lis, err := net.Listen("tcp", GrpcServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen: %v", err)
	}

	s := handler.Profile{}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	profile.RegisterProfileServer(grpcServer, &s)

	log.Info().Msgf("start gRPC server at %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msgf("failed to serve: %s", err)
	}
}
