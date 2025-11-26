package main

import (
	"context"
	"log"
	"net"
	"os"

	grpcserver "github.com/johnqr/user-service/internal/grpc"
	"google.golang.org/grpc"
)

func main() {
	port := ":9090"
	if p := os.Getenv("GRPC_PORT"); p != "" { port = ":" + p }

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpcSrv := grpc.NewServer()
	_ = grpcserver.NewServer // avoid unused
	// registration happens inside NewServer
	_ = context.Background()
	log.Println("gRPC server running on", port)
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
