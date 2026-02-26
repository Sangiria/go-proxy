package handlers

import (
	"context"
	"core/api"
	"core/file"
	"core/links"
	"core/models"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeService struct {

}

func GetNodesHandler() {

}

//TODO: better error handling and id generation

func (n *NodeService) AddNodeHandler(ctx context.Context, message *api.AddNodeRequest) (*api.AddNodeResponse, error) {
	//get state variable
	state, err := file.LoadState()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something went wrong while loading state: %s", err)
	}
	
	//parcing url and if https fetching nodes
	switch(message.Source) {
	case api.Source_Manual:
		if err := state.AddNodeFromURL(message.Url, &models.Source{
			Type: message.Source.String(),
		}); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

	case api.Source_Subscription:
		sub_key, err := state.AddSubscriptionFromURL(message.Url)
		if err != nil {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		links, err := links.FetchVLESSLinks(message.Url)

		for _, link := range links{
			if err := state.AddNodeFromURL(link, &models.Source{
				Type: message.Source.String(),
				SubscriptionID: sub_key,
			}); err != nil {
				continue
			}
		}
	}

	//update file
	t0 := time.Now()
	if err = file.SaveState(state); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	dt := time.Since(t0)
	info, err := os.Stat("./state/state.json")
	if err != nil {
		log.Fatal(err)
	}

	size := info.Size()

	fmt.Printf("Size: %d bytes (%.2f KiB)\n", size, float64(size)/1024)
	fmt.Printf("Time: %v\n", dt)

	speed := float64(size) / dt.Seconds() / 1024
	fmt.Printf("Speed: %.2f KiB/s\n", speed)

	return &api.AddNodeResponse{}, nil
}