package main

import (
	"core/api"
	nodeservice "core/api/node_service"
	proxyservice "core/api/proxy_service"
	"core/internal/file"
	"core/internal/manager"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
  		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	state, err := file.LoadState()
    if err != nil {
        log.Fatalf("failed to load state: %v", err)
    }

	manager := manager.NewManager(state)

	node_service := nodeservice.NewNodeService(manager)
	proxy_service := proxyservice.NewProxyService(manager)

	api.RegisterNodeServiceServer(grpcServer, node_service)
	api.RegisterProxyServiceServer(grpcServer, proxy_service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}