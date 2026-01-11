package main

import (
	"bufio"
	"core/handlers"
	"fmt"
	"log"
	"os"
)

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

	answer, err := handlers.AddNodesHandler(line)
	if err != nil {
		fmt.Printf("%s:%s",answer, err)
	}

	fmt.Printf(answer)
}