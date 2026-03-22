package nodeservice

import (
	"core/api"
	"core/internal/file"
	"core/internal/links"
	"core/models"
	"errors"
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

func (n *NodeService) findSubscription(id string) *models.Subscription {
	sub, ok := n.state.Subscriptions[id]
	if !ok {
		return nil
	}
	return sub
}

func (n *NodeService) findNode(message *api.Id) *models.Node {
	var target *models.Node

    if message.SourceId != nil {
        sub := n.findSubscription(*message.SourceId)
        if sub == nil {
            return nil
        }
        target = sub.Nodes[message.Id]
    } else {
        target = n.state.Manual[message.Id]
    }

	return target
}

func (n *NodeService) updateSubscriptionNodes(sub *models.Subscription) (*api.Nodes, error){
	node_links, err := links.FetchVLESSLinks(sub.URL)
	if err != nil {
		return nil, errors.New("error getting nodes")
	}

	var (
		added = make([]*api.Node, len(node_links))
		nodes = make(map[string]*models.Node, len(node_links))
		nodes_order = make([]string, len(node_links))
	)

	for id, link := range node_links{
		node_key := links.GenerateID(link)

		node, err := links.ParseURLToNode(link)
		if err != nil {
			return nil, err
		}

		nodes[node_key] = node
		nodes_order[id] = node_key
		added[id] = mapToApiNode(node_key, node)
	}

	sub.Nodes, sub.NodeOrder = nodes, nodes_order
	
	return &api.Nodes{
		Nodes: added,
	}, nil
}