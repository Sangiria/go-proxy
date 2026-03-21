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
		items = make([]*api.Id, len(n.state.ItemsOrder))
		manual = make([]*api.Node, len(n.state.Manual))
		sub = make([]*api.Subscription, len(n.state.Subscriptions))
		node_id = 0
		sub_id = 0
	)

	for item_id, item_key := range n.state.ItemsOrder {
		_, node_exists := n.state.Manual[item_key]
		if node_exists {
			manual[node_id] = mapToApiNode(item_key, n.state.Manual[item_key])
			node_id++
		}

		_, sub_exists := n.state.Subscriptions[item_key]
		if sub_exists {
			var nodes = make([]*api.Node, len(n.state.Subscriptions[item_key].NodeOrder))

			for snode_id, snode_key := range n.state.Subscriptions[item_key].NodeOrder {
				node := n.state.Subscriptions[item_key].Nodes[snode_key]
				nodes[snode_id] = mapToApiNode(snode_key, node)
			}

			sub[sub_id] = &api.Subscription{
				Id: item_key,
				Name: n.state.Subscriptions[item_key].Name,
				Nodes: nodes,
			}

			sub_id++
		}

		items[item_id] = &api.Id{
			Id: item_key,
		}
	}

	return &api.State{
		Manual: manual,
		Subscription: sub,
		Order: items,
	}, nil
}