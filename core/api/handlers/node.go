package handlers

import (
	"context"
	"core/api"
	"core/internal/file"
	"core/internal/links"
	"core/internal/service"
	"core/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func mapToApiNode(id string, node *models.Node) *api.Node {
    return &api.Node{
        Id:        	id,
        Type:      	node.Parsed.Type,
        Name:      	node.Name,
        Address:   	node.Parsed.Address,
        Port:      	int32(node.Parsed.Port),
        Transport: 	node.Parsed.Transport,
        Security: 	node.Parsed.Security,
    }
}

func mapToApiNodeForm(node *models.Node) *api.NodeForm {
	port := int32(node.Parsed.Port)
	extra := string(node.Parsed.XHTTPExtra)

	return &api.NodeForm{
		Name: &node.Name,
		Address: &node.Parsed.Address,
		Port: &port,
		Uuid: &node.Parsed.UUID,
		Transport: &node.Parsed.Transport,
		Security: &node.Parsed.Security,
		Sni: &node.Parsed.Sni,
		Fp: &node.Parsed.Fp,
		Pbk: &node.Parsed.Pbk,
		Sid: &node.Parsed.Sid,
		Mode: &node.Parsed.XHTTPMode,
		Extra: &extra,
	}
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

func (n *NodeService) EditNode(ctx context.Context, message *api.NodeForm) (*api.Null, error) {
    empty := &api.NodeForm{Id: message.Id, SourceId: message.SourceId}
    if proto.Equal(message, empty) {
        return nil, status.Errorf(codes.InvalidArgument, "nothing to update")
    }

    node := n.FindNode(&api.Id{Id: message.Id, SourceId: message.SourceId})

    if err := service.UpdateNodeFromForm(node, message); err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "update failed: %v", err)
    }

    if err := file.SaveState(n.state); err != nil {
        return nil, status.Errorf(codes.Internal, "save error: %v", err)
    }
    
    return &api.Null{}, nil
}

func (n *NodeService) GetNode(ctx context.Context, message *api.Id) (*api.NodeForm, error) {
	return mapToApiNodeForm(n.FindNode(message)), nil
}