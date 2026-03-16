package handlers

import (
	"context"
	"core/api"
	"core/internal/file"
	"core/models"
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

func (n *NodeService) FindSubscription(id string) *models.Subscription {
	sub, ok := n.state.Subscriptions[id]
	if !ok {
		return nil
	}
	return sub
}

func (n *NodeService) FindNode(message *api.Id) *models.Node {
	var target *models.Node

    if message.SourceId != nil {
        sub := n.FindSubscription(*message.SourceId)
        if sub == nil {
            return nil
        }
        target = sub.Nodes[message.Id]
    } else {
        target = n.state.Manual[message.Id]
    }

	return target
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
			nodes[node_id] = mapToApiNode(node_key, node)
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