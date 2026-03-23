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
	status_chan 	chan api.ProxyState
}

func NewProxyService(manager *manager.Manager) *ProxyService {
	return &ProxyService{
		mg: manager,
		status_chan: make(chan api.ProxyState, 10),
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

        if strings.Contains(line, "initial connection") {
            p.status_chan <- api.ProxyState_CONNECTED
        } else if strings.Contains(line, "address already in use") {
            p.status_chan <- api.ProxyState_ERROR
            fmt.Println("!!! PORT ALREADY IN USE !!!")
        }
    }

	// p.status_chan <- api.ProxyState_DISCONNECTED
}