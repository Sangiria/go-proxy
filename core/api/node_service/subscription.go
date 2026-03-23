package nodeservice

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
	sub := n.mg.FindSubscription(message.Id)
	data, err := links.FetchSubscriptionNodes(sub.URL)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	n.mg.Mu.Lock()
    defer n.mg.Mu.Unlock()

	sub.Nodes, sub.NodeOrder = data.Nodes, data.Order

	if err := file.SaveState(n.mg.State); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.Nodes{
		Nodes: data.Added,
	}, nil
}

func (n *NodeService) AddSubscription(ctx context.Context, message *api.Url) (*api.Subscription, error) {
	n.mg.Mu.Lock()
    defer n.mg.Mu.Unlock()

	sub_key := links.GenerateID(message.Url)

	_, found := n.mg.State.Subscriptions[sub_key]
	if found {
		return nil, status.Errorf(codes.AlreadyExists, "subscription already exists")
	}

	name, err := url.Parse(message.Url)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sub := &models.Subscription{
		Name: name.Host,
		URL: message.Url,
		Nodes: make(map[string]*models.Node, 0),
		NodeOrder: make([]string, 0),
	}

	data, err := links.FetchSubscriptionNodes(sub.URL)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}
	sub.Nodes, sub.NodeOrder = data.Nodes, data.Order

	n.mg.State.Subscriptions[sub_key] = sub
	n.mg.State.ItemsOrder = append(n.mg.State.ItemsOrder, sub_key)

	if err := file.SaveState(n.mg.State); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.Subscription{
		Id: sub_key,
		Name: name.Host,
		Nodes: &api.Nodes{
			Nodes: data.Added,
		},
	}, nil
}

func (n *NodeService) EditSubscription(ctx context.Context, message *api.SubscriptionForm) (*api.Null, error) {
	n.mg.Mu.Lock()
    defer n.mg.Mu.Unlock()
	
	empty := &api.SubscriptionForm{Id: message.Id}

	if proto.Equal(message, empty) {
		return nil, status.Errorf(codes.InvalidArgument, "nothing to update")
	}

	if err := service.UpdateSubscriptionFromForm(n.mg.State.Subscriptions[message.Id], message); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "update failed: %v", err)
	}

	if err := file.SaveState(n.mg.State); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	return &api.Null{}, nil
}

func (n *NodeService) GetSubscription(ctx context.Context, message *api.Id) (*api.SubscriptionForm, error) {
	sub := n.mg.FindSubscription(message.Id)
	return &api.SubscriptionForm{
		Name: &sub.Name,
		Url: &sub.URL,
	}, nil
}

func (n *NodeService) DeleteSubscription(ctx context.Context, message *api.Id) (*api.Null, error) {
	n.mg.Mu.Lock()
    defer n.mg.Mu.Unlock()

	_, ok := n.mg.State.Subscriptions[message.Id].Nodes[n.mg.State.ActiveNodeId]
	if ok {
		return nil, status.Errorf(codes.PermissionDenied, "this subscription has active node")
	}

	for id, sub_key := range n.mg.State.ItemsOrder {
		if sub_key == message.Id {
			n.mg.State.ItemsOrder = append(n.mg.State.ItemsOrder[:id], n.mg.State.ItemsOrder[id+1:]...)
			delete(n.mg.State.Subscriptions, sub_key)
		}
	}

	if err := file.SaveState(n.mg.State); err != nil {
        return nil, status.Errorf(codes.Internal, "save error: %v", err)
    }

	return &api.Null{}, nil
}