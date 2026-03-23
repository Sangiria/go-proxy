package proxyservice

import (
	"context"
	"core/api"
	"core/models"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *ProxyService) StartProxy(ctx context.Context, message *api.Id) (*api.Null, error) {
    node := p.mg.FindNode(message)
    if node == nil {
        return nil, status.Error(codes.NotFound, "node not found")
    }

    cfg := models.NewConfig(&node.Parsed)
    configJson, err := json.MarshalIndent(cfg, "", "  ")
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to marshal config: %v", err)
    }

    executablePath := "../bin/xray" 
    cmd := exec.Command(executablePath, "run", "-config", "stdin:")
    
    stdin, err := cmd.StdinPipe()
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to create stdin pipe: %v", err)
    }

    stderr, _ := cmd.StderrPipe()
    cmd.Stdout = cmd.Stderr

    if err := cmd.Start(); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to start xray: %v", err)
    }

    p.mg.SetActiveNode(message.Id)

    go func() {
        defer stdin.Close()
        _, err := stdin.Write(configJson)
        if err != nil {
            fmt.Printf("failed to write to stdin: %v\n", err)
        }
    }()

    p.setCurrentCmd(cmd)
    go p.monitorXrayLogs(stderr)
    
    return &api.Null{}, nil
}

func (p *ProxyService) StopProxy(ctx context.Context, req *api.Null) (*api.Null, error) {
    p.mu.Lock()
    defer p.mu.Unlock()

    if p.current_cmd == nil || p.current_cmd.Process == nil {
        return &api.Null{}, nil
    }

    if err := p.current_cmd.Process.Signal(os.Interrupt); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to stop xray process: %v", err)
    }

    p.current_cmd.Wait()
    p.current_cmd = nil

    p.updateStatus(&api.ProxyStatus{State: api.ProxyState_DISCONNECTED})
    p.mg.ClearActiveNode()

    fmt.Println("Proxy stopped and system settings cleared")
    return &api.Null{}, nil
}

func (p *ProxyService) SubscribeStatus(req *api.Null, stream api.ProxyService_SubscribeStatusServer) error {
    state := api.ProxyState_DISCONNECTED
    if p.current_cmd != nil {
        state = api.ProxyState_CONNECTED
    }
    
    stream.Send(&api.ProxyStatus{State: state})

    for {
        select {
        case <-stream.Context().Done():
            return stream.Context().Err()

        case status, ok := <-p.status_chan:
            if !ok {
                return nil
            }

            if err := stream.Send(status); err != nil {
                return err
            }
        }
    }
}