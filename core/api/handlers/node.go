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

func (n *NodeService) AddNode(ctx context.Context, message *api.AddNodeRequest) (*api.AddNodeResponse, error) {
	node_key := links.GenerateID(message.Url)
	
	_, found := n.state.Nodes[node_key]
	if found {
		return nil, status.Errorf(codes.AlreadyExists, "node already exist")
	}

	node, err := links.ParseURLToNode(message.Url, &models.Source{
		Type: models.SourceManual,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	n.state.Nodes[node_key] = node

	if err = file.SaveState(n.state); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddNodeResponse{}, nil
}

//TODO: вынести получение нод подписок и их добавление в state в отдельный rpc метод
//TODO: прибавлять айди подписки к ключу ноды при генерации айди
//TODO: возвращать при ответе количество добавленных нодов и количество всего полученных нодов из ссылки

func (n *NodeService) AddSubscription(ctx context.Context, message *api.AddNodeRequest) (*api.AddNodeResponse, error) {
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

	node_links, err := links.FetchVLESSLinks(message.Url)

	for _, link := range node_links{
		node_key := links.GenerateID(link)
		if _, found := n.state.Nodes[node_key]; found {
			continue
		}

		node, err := links.ParseURLToNode(link, &models.Source{
			Type: models.SourceSubscription,
			SubscriptionID: sub_key,
		})

		if err != nil {
			continue
		}

		n.state.Nodes[node_key] = node
	}

	if err = file.SaveState(n.state); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddNodeResponse{}, nil
}