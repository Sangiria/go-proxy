package main

import (
	"core/api"
	nodeservice "core/api/node_service"
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
	node_service, err := nodeservice.NewNodeService()
	if err != nil {
		log.Fatalf("failed to create node service: %v", err)
	}

	api.RegisterNodeServiceServer(grpcServer, node_service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}