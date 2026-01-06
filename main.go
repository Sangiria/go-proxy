package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/proxy"
)

type LogLevel struct {
  	LogLevel string `json:"loglevel"`
}

type Inbound struct {
	Tag       		string      	`json:"tag"`
  	Listen     	 	string      	`json:"listen"`
  	Port      		uint16      	`json:"port"`
  	Protocol    	string      	`json:"protocol"`
  	InboundSettings InboundSettings `json:"settings"`
}

type InboundSettings struct {
  	Auth  string  	`json:"auth"`
  	Udp   bool  	`json:"udp"`
}

type Outbound struct {
  	Tag            string         `json:"tag"`
  	Protocol       string         `json:"protocol"`
  	Settings       VlessSettings  `json:"settings"`
  	StreamSettings StreamSettings `json:"streamSettings"`
}

type VlessSettings struct {
  	VNext []VNext `json:"vnext"`
}

type VNext struct {
  	Address string		`json:"address"`
  	Port    uint16		`json:"port"`
  	Users   []VlessUser	`json:"users"`
}

type VlessUser struct {
  	ID         string `json:"id"`
  	Encryption string `json:"encryption"`
  	Flow       string `json:"flow,omitempty"`
}

type StreamSettings struct {
  	Network         string          `json:"network"`
  	Security        string          `json:"security"`
  	RealitySettings RealitySettings `json:"realitySettings"`
}

type RealitySettings struct {
  	ServerName  string `json:"serverName"`
  	Fingerprint string `json:"fingerprint"`
  	PublicKey   string `json:"publicKey"`
  	ShortID     string `json:"shortId"`
}

type Config struct {
  	LogLevel    LogLevel	`json:"log"`
  	Inbounds   []Inbound	`json:"inbounds"`
  	Outbounds  []Outbound	`json:"outbounds"`
}


func NewConfig(s string) (*Config, error) {
	s = strings.TrimSpace(s)

	u, err := url.Parse(s)
  	if err != nil {
    	return nil, fmt.Errorf("invalid url: %w", err)
  	}
  	q_u := u.Query()

  	if err = ValidateVlessLink(u, q_u); err != nil {
    	return nil, fmt.Errorf("invalid url: %w", err)
  	}

  	port, _ := strconv.ParseUint(u.Port(), 10, 16)
  
  	return &Config{
		LogLevel: LogLevel{
			LogLevel: "warning",
    	},
    	Inbounds: []Inbound{
			{
        		Tag: "socks",
        		Listen: "127.0.0.1",
        		Port: 10808,
				Protocol: "socks",
        		InboundSettings: InboundSettings{
          			Auth: "noauth",
          			Udp: false,
        		},
      		},
    	},
    	Outbounds: []Outbound{
			{
        		Tag: "proxy",
        		Protocol: "vless",
        		Settings: VlessSettings{
          			VNext: []VNext{
            			{
              				Address: u.Hostname(),
              				Port: uint16(port),
              				Users: []VlessUser{
                				{
                  					ID: u.User.Username(),
                  					Encryption: "none",
                  					Flow: q_u.Get("flow"),
                				},
              				},
            			},
          			},
        		},
        		StreamSettings: StreamSettings{
          			Network: q_u.Get("type"),
          			Security: q_u.Get("security"),
          			RealitySettings: RealitySettings{
            			ServerName: q_u.Get("sni"),
            			Fingerprint: q_u.Get("fp"),
            			PublicKey: q_u.Get("pbk"),
            			ShortID: strings.TrimRight(q_u.Get("sid"), "#"),
          			},
        		},
      		},
    	},
  	}, nil
}

func ValidateVlessLink(u *url.URL, q_u url.Values) error {
	if u.Scheme != "vless" || u.User == nil || u.User.Username() == "" {
    	return fmt.Errorf("not a vless url")
  	}

  	if err := uuid.Validate(u.User.Username()); err != nil {
    	return fmt.Errorf("invalid uuid")
	}

  	if u.Hostname() == "" || u.Port() == "" {
    	return fmt.Errorf("missing host or port")
  	}

  	port, err := strconv.Atoi(u.Port())
  	if err != nil || port < 1 || port > 65535 {
		  return fmt.Errorf("invalid port")
		}
		
		if q_u.Get("security") != "reality" || q_u.Get("type") != "tcp" {
    	return fmt.Errorf("unsupported")
	}
	
	required := []string{"sni", "fp", "pbk", "sid"}
	for _, k := range required {
		if strings.TrimSpace(q_u.Get(k)) == "" {
			return fmt.Errorf("missing %s", k)
    	}
	}
	
	return nil
}

var xrayPath = "./bin/xray"
var configPath = "config.json"

func generateTrafficViaSocks(socksAddr string) error {
    dialer, err := proxy.SOCKS5("tcp", socksAddr, nil, proxy.Direct)
    if err != nil {
        return fmt.Errorf("socks dialer: %w", err)
    }

    tr := &http.Transport{
        DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
            return dialer.Dial(network, addr)
        },
        TLSClientConfig: &tls.Config{MinVersion: tls.VersionTLS12},
    }

    client := &http.Client{
        Transport: tr,
        Timeout:   8 * time.Second,
    }

    resp, err := client.Get("https://google.com")
    if err != nil {
        return fmt.Errorf("http via socks: %w", err)
    }
    defer resp.Body.Close()
    _, _ = io.ReadAll(io.LimitReader(resp.Body, 256))
    return nil
}

func main() {
	fmt.Print("Enter url: ")
  	reader := bufio.NewReader(os.Stdin)
  	line, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("read line: %s\n", line)

	conf, err := NewConfig(line)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("config.json", data, 0644)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, xrayPath, "run", "-config", configPath)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(fmt.Errorf("xray start error: %w", err))
	}

	go stream("XRAY-OUT", stdout)
	go stream("XRAY-ERR", stderr)

	if err := waitTCP("127.0.0.1:10808", 3*time.Second); err != nil {
		_ = cmd.Process.Kill()
		panic(fmt.Errorf("xray not ready: %w", err))
	}

	fmt.Println("✅ xray started, inbound is ready on 127.0.0.1:10808")

	go func() {
    	time.Sleep(300 * time.Millisecond)

    	err := generateTrafficViaSocks("127.0.0.1:10808")
    	if err != nil {
        	fmt.Println("[TRAFFIC] error:", err)
    	} else {
        	fmt.Println("[TRAFFIC] ok")
    	}
	}()

	time.Sleep(20 * time.Second)

	cancel()

	_ = cmd.Wait()
	fmt.Println("🛑 xray stopped")
}

func stream(prefix string, r interface{ Read([]byte) (int, error) }) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Printf("[%s] %s\n", prefix, sc.Text())
	}
}

func waitTCP(addr string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", addr, 250*time.Millisecond)
		if err == nil {
			_ = conn.Close()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("port %s not open within %s", addr, timeout)
}