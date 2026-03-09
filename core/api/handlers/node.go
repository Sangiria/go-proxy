package handlers

import (
	"context"
	"core/api"
	"core/file"
	"core/links"
	"core/models"
	"net/url"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeService struct {
	api.UnimplementedNodeServiceServer
	state *file.State
}

func NewNodeService() (*NodeService, error) {
    state, err := file.LoadState()
    if err != nil {
        return nil, err
    }
    return &NodeService{state: state}, nil
}

func (n *NodeService) storeNode(url string, source *models.Source) error {
	node_key := links.GenerateID(url)
	
	_, found := n.state.Nodes[node_key]
	if found {
		return status.Errorf(codes.AlreadyExists, "node already exist")
	}

	node, err := links.ParseURLToNode(url, source)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	n.state.Nodes[node_key] = node

	return nil
}

func (n *NodeService) GetNodes(ctx context.Context, message *api.Null) (*api.Nodes, error) {
	nodes := make([]*api.Node, 0, len(n.state.Nodes))

	for id, node := range n.state.Nodes {
		nodes = append(nodes, &api.Node{
			Id: id,
			Type: node.Parsed.Type,
			Name: node.Name,
			Address: node.Parsed.Address,
			Port: int32(node.Parsed.Port),
			Transport: node.Parsed.Transport,
			Tls: node.Parsed.Security,
			Source: string(node.Source.Type),
		})
	}

	return &api.Nodes{
		Nodes: nodes,
	}, nil
}

func (n *NodeService) AddNode(ctx context.Context, message *api.Url) (*api.AddedNodesCount, error) {
	if err := n.storeNode(message.Url, &models.Source{
		Type: models.SourceManual,
	}); err != nil {
		return nil, err
	}

	if err := file.SaveState(n.state); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddedNodesCount{
		Added: 1,
		Received: 1,
	}, nil
}

func (n *NodeService) FetchSubscriptionNodes(ctx context.Context, message *api.SubscriptionId) (*api.AddedNodesCount, error) {
	url := n.state.Subscriptions[message.Id].URL
	node_links, err := links.FetchVLESSLinks(url)

	var (
		added = 0
		recieved = len(node_links)
	)

	for _, link := range node_links{
		if err := n.storeNode(link, &models.Source{
			Type: models.SourceSubscription,
			SubscriptionID: message.Id,
		}); err != nil {
			continue
		}

		added++
	}

	if err = file.SaveState(n.state); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddedNodesCount{
		Added: int64(added),
		Received: int64(recieved),
	}, nil
}

func (n *NodeService) AddSubscription(ctx context.Context, message *api.Url) (*api.SubscriptionId, error) {
	sub_key := links.GenerateID(message.Url)

	_, found := n.state.Subscriptions[sub_key]
	if found {
		return nil, status.Errorf(codes.AlreadyExists, "subscription already exists")
	}

	name, err := url.Parse(message.Url)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	n.state.Subscriptions[sub_key] = &models.Subscription{
		Name: name.Host,
		URL: message.Url,
	}

	return &api.SubscriptionId{
		Id: sub_key,
	}, nil
}