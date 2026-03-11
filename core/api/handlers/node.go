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

func mapToApiNode(id string, node *models.Node) *api.Node {
    return &api.Node{
        Id:        id,
        Type:      node.Parsed.Type,
        Name:      node.Name,
        Address:   node.Parsed.Address,
        Port:      int32(node.Parsed.Port),
        Transport: node.Parsed.Transport,
        Tls:       node.Parsed.Security,
    }
}

func NewNodeService() (*NodeService, error) {
    state, err := file.LoadState()
    if err != nil {
        return nil, err
    }
    return &NodeService{state: state}, nil
}

func (n *NodeService) AddManual(ctx context.Context, message *api.Url) (*api.Node, error) {
	node_key := links.GenerateID(message.Url)

	_, found := n.state.Manual[node_key]
	if found {
		return nil, status.Errorf(codes.AlreadyExists, "manual node already exist")
	}

	node, err := links.ParseURLToNode(message.Url)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	n.state.Manual[node_key] = node

	if err := file.SaveState(n.state); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return mapToApiNode(node_key, node), nil
}

func (n *NodeService) AddSubscription(ctx context.Context, message *api.Url) (*api.Subscription, error) {
	sub_key := links.GenerateID(message.Url)

	_, found := n.state.Subscriptions[sub_key]
	if found {
		return nil, status.Errorf(codes.AlreadyExists, "subscription already exists")
	}

	name, err := url.Parse(message.Url)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	node_links, err := links.FetchVLESSLinks(message.Url)
	if err != nil {
		return nil, status.Errorf(codes.Canceled, "error getting subscribtion")
	}

	n.state.Subscriptions[sub_key] = &models.Subscription{
		Name: name.Host,
		URL: message.Url,
		Nodes: make(map[string]models.Node, len(node_links)),
	}

	var nodes = make([]*api.Node, 0, len(node_links))

	for id, link := range node_links{
		node_key := links.GenerateID(link)

		node, err := links.ParseURLToNode(link)
		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}

		n.state.Subscriptions[sub_key].Nodes[node_key] = *node
		nodes[id] = mapToApiNode(node_key, node)
	}

	if err := file.SaveState(n.state); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.Subscription{
		Id: sub_key,
		Name: name.Host,
		Nodes: nodes,
	}, nil
}