package nodeservice

import (
	"core/api"
	"core/internal/manager"
)

type NodeService struct {
	api.UnimplementedNodeServiceServer
	mg *manager.Manager
}

func NewNodeService(manager *manager.Manager) *NodeService {
	return &NodeService{mg: manager}
}