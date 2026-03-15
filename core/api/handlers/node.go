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

func (n *NodeService) GetFullState(ctx context.Context, message *api.Null) (*api.State, error) {
	if n.state.Manual == nil && n.state.Subscriptions == nil {
		return &api.State{}, nil
	}

	var (
		manual = make([]*api.Node, len(n.state.ManualOrder))
		sub = make([]*api.Subscription, len(n.state.SubscriptionOrder))
	)

	for node_id, node_key := range n.state.ManualOrder {
		manual[node_id] = mapToApiNode(node_key, n.state.Manual[node_key])
	}

	for sub_id, sub_key := range n.state.SubscriptionOrder {
		var nodes = make([]*api.Node, len(n.state.Subscriptions[sub_key].NodeOrder))

		for node_id, node_key := range n.state.Subscriptions[sub_key].NodeOrder {
			node := n.state.Subscriptions[sub_key].Nodes[node_key]
			nodes[node_id] = mapToApiNode(node_key, &node)
		}

		sub[sub_id] = &api.Subscription{
			Id: sub_key,
			Name: n.state.Subscriptions[sub_key].Name,
			Nodes: nodes,
		}
	}

	return &api.State{
		Manual: manual,
		Subscription: sub,
	}, nil
}

func (n *NodeService) AddNode(ctx context.Context, message *api.Url) (*api.Node, error) {
	node_key := links.GenerateID(message.Url)

	_, found := n.state.Manual[node_key]
	if found {
		return nil, status.Errorf(codes.AlreadyExists, "manual node already exists")
	}

	node, err := links.ParseURLToNode(message.Url)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	n.state.Manual[node_key] = node
	n.state.ManualOrder = append(n.state.ManualOrder, node_key)

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
		NodeOrder: make([]string, 0),
	}
	n.state.SubscriptionOrder = append(n.state.SubscriptionOrder, sub_key)

	var nodes = make([]*api.Node, len(node_links))

	for id, link := range node_links{
		node_key := links.GenerateID(link)

		node, err := links.ParseURLToNode(link)
		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}

		n.state.Subscriptions[sub_key].Nodes[node_key] = *node
		n.state.Subscriptions[sub_key].NodeOrder = append(n.state.Subscriptions[sub_key].NodeOrder, node_key)

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