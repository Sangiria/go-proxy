package main

import (
	"core/api"
	nodeservice "core/api/node_service"
	proxyservice "core/api/proxy_service"
	"core/internal/file"
	"core/internal/manager"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Println("gRPC server started on :3333")
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-stop
	log.Println("\nShutting down gracefully...")

	if err := file.SaveState(manager.State); err != nil {
		log.Printf("Error saving state during shutdown: %v", err)
	} else {
		log.Println("State saved successfully.")
	}

	stopped := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		log.Println("Server stopped.")
	case <-time.After(5 * time.Second):
		log.Println("Shutdown timed out, forcing stop.")
		grpcServer.Stop()
	}
}