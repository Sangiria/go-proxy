package handlers

import (
	"context"
	"core/api"
	"core/internal/file"
	"core/internal/links"
	"core/internal/service"
	"core/models"
	"net/url"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func (n *NodeService) UpdateSubscription(ctx context.Context, message *api.Id) (*api.Nodes, error)  {
	sub := n.FindSubscription(message.Id)

	nodes, err := n.UpdateSubscriptionNodes(sub)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}
	return nodes, nil
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

	n.state.Subscriptions[sub_key] = &models.Subscription{
		Name: name.Host,
		URL: message.Url,
		Nodes: make(map[string]*models.Node, 0),
		NodeOrder: make([]string, 0),
	}

	nodes, err := n.UpdateSubscriptionNodes(n.state.Subscriptions[sub_key])
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	n.state.ItemsOrder = append(n.state.ItemsOrder, sub_key)

	if err := file.SaveState(n.state); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.Subscription{
		Id: sub_key,
		Name: name.Host,
		Nodes: nodes,
	}, nil
}

func (n *NodeService) EditSubscription(ctx context.Context, message *api.SubscriptionForm) (*api.Null, error) {
	empty := &api.SubscriptionForm{Id: message.Id}

	if proto.Equal(message, empty) {
		return nil, status.Errorf(codes.InvalidArgument, "nothing to update")
	}

	if err := service.UpdateSubscriptionFromForm(n.state.Subscriptions[message.Id], message); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "update failed: %v", err)
	}

	if err := file.SaveState(n.state); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	return &api.Null{}, nil
}

func (n *NodeService) GetSubscription(ctx context.Context, message *api.Id) (*api.SubscriptionForm, error) {
	sub := n.FindSubscription(message.Id)

	return &api.SubscriptionForm{
		Name: &sub.Name,
		Url: &sub.URL,
	}, nil
}

func (n *NodeService) DeleteSubscription(ctx context.Context, message *api.Id) (*api.Null, error) {
	_, ok := n.state.Subscriptions[message.Id].Nodes[n.state.ActiveNodeId]
	if ok {
		return nil, status.Errorf(codes.PermissionDenied, "this subscription has active node")
	}

	for id, sub_key := range n.state.ItemsOrder {
		if sub_key == message.Id {
			n.state.ItemsOrder = append(n.state.ItemsOrder[:id], n.state.ItemsOrder[id+1:]...)
			delete(n.state.Subscriptions, sub_key)
		}
	}

	return &api.Null{}, nil
}