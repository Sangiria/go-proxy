package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

// var xrayPath = "./bin/xray"
// var configPath = "config.json"

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
  		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}