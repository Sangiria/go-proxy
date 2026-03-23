package manager

import (
	"core/api"
	"core/internal/file"
	"core/models"
	"sync"
)

func MapToApiNode(id string, node *models.Node) *api.Node {
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

type Manager struct {
	State *file.State
	Mu    sync.RWMutex
}

func NewManager(s *file.State) *Manager {
	return &Manager{State: s}
}
func (m *Manager) GetActiveNodeID() string {
	m.Mu.RLock()
    defer m.Mu.RUnlock()

	return m.State.ActiveNodeId
}

func (m *Manager) SetActiveNode(id string) {
	m.Mu.RLock()
    defer m.Mu.RUnlock()
	m.State.ActiveNodeId = id

	_ = file.SaveState(m.State)
}

func (m *Manager) ClearActiveNode() {
	m.Mu.RLock()
    defer m.Mu.RUnlock()
	m.State.ActiveNodeId = ""

	_ = file.SaveState(m.State)
}

func (m *Manager) FindNode(message *api.Id) *models.Node {
	m.Mu.RLock()
    defer m.Mu.RUnlock()
	
	var target *models.Node
    if message.SourceId != nil {
        sub := m.FindSubscription(*message.SourceId)
        if sub == nil {
            return nil
        }
        target = sub.Nodes[message.Id]
    } else {
        target = m.State.Manual[message.Id]
    }

	return target
}

func (m *Manager) FindSubscription(id string) *models.Subscription {
	m.Mu.RLock()
    defer m.Mu.RUnlock()

	sub, ok := m.State.Subscriptions[id]
	if !ok {
		return nil
	}
	return sub
}