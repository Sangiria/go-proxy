package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

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

func main() {
	fmt.Print("Enter url: ")
  	reader := bufio.NewReader(os.Stdin)
  	line, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("read line: %s\n", line)

	// conf, err := NewConfig(line)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, err := json.MarshalIndent(conf, "", "  ")
	// if err != nil {
	// 	panic(err)
	// }

	// err = os.WriteFile("config.json", data, 0644)
}