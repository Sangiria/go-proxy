package proxyservice

import (
	"bufio"
	"core/api"
	"core/internal/manager"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type ProxyService struct {
	api.UnimplementedProxyServiceServer
	mg 				*manager.Manager
	current_cmd 	*exec.Cmd
    mu      		sync.Mutex
	status_chan     chan *api.ProxyStatus
}

func NewProxyService(manager *manager.Manager) *ProxyService {
	return &ProxyService{
		mg: manager,
		status_chan: make(chan *api.ProxyStatus, 10),
	}
}

func (p *ProxyService) updateStatus(state api.ProxyState, msg string) {
    activeID := p.mg.GetActiveNodeID()

    if state == api.ProxyState_DISCONNECTED || state == api.ProxyState_ERROR {
        activeID = ""
    }

    status := &api.ProxyStatus{
        State:        state,
        Message:      msg,
        ActiveNodeId: activeID,
    }

    select {
    case p.status_chan <- status:
    default:
    }
}

func (p *ProxyService) setCurrentCmd(cmd *exec.Cmd) {
    p.mu.Lock()
    defer p.mu.Unlock()
    
    if p.current_cmd != nil && p.current_cmd.Process != nil {
        // p.mg.SetSystemProxy(false, 0)

        _ = p.current_cmd.Process.Signal(os.Interrupt)
        
        done := make(chan error, 1)
        go func() {
            done <- p.current_cmd.Wait()
        }()

        select {
        case <-done:
            fmt.Println("Xray exited gracefully")
        case <-time.After(200 * time.Millisecond):
            p.current_cmd.Process.Kill()
        }
    }
    p.current_cmd = cmd
}

func (p *ProxyService) monitorXrayLogs(reader io.ReadCloser) {
    defer reader.Close()
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("[XRAY LOG]: %s\n", line)

        switch {
        case strings.Contains(line, "started"):
            p.updateStatus(api.ProxyState_CONNECTED, "")
        case strings.Contains(line, "address already in use"):
            p.updateStatus(api.ProxyState_ERROR, "port is already in use")
        case strings.Contains(line, "failed to handler mux client") && strings.Contains(line, "io: read/write on closed pipe"):
            p.updateStatus(api.ProxyState_ERROR, "protocol error, check key settings")
        }
    }

    p.mg.ClearActiveNode()
    p.updateStatus(api.ProxyState_DISCONNECTED, "")
}