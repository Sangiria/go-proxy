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
		Nodes: make(map[string]*models.Node, len(node_links)),
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

		n.state.Subscriptions[sub_key].Nodes[node_key] = node
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