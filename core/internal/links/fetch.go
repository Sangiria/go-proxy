package links

import (
	"core/api"
	"core/internal/manager"
	"core/models"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type UpdateData struct {
    Nodes     	map[string]*models.Node
    Order     	[]string
    Added  		[]*api.Node
}

func FetchSubscriptionNodes (url string) (*UpdateData, error){
	node_links, err := FetchVLESSLinks(url)
	if err != nil {
		return nil, errors.New("error getting nodes")
	}

	data := &UpdateData{
        Nodes:    	make(map[string]*models.Node),
        Order:    	make([]string, len(node_links)),
        Added: 		make([]*api.Node, len(node_links)),
    }

	for id, link := range node_links{
		node_key := GenerateID(link)

		node, err := ParseURLToNode(link)
		if err != nil {
			return nil, err
		}

		data.Nodes[node_key] = node
		data.Order[id] = node_key
		data.Added[id] = manager.MapToApiNode(node_key, node)
	}

	return data, nil
}


func FetchVLESSLinks(u string) ([]string, error){
	response, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %v\n", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: Received non-OK status code: %d\n", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v\n", err)
	}

	decoded, err := decodeBase64(string(body))
	if err != nil {
		return nil, fmt.Errorf("Invalid body contents: %v\n", err)
	}
	links := strings.Split(string(decoded), "\n")

	return links, nil
}

func decodeBase64(s string) ([]byte, error) {
	encodings := []*base64.Encoding{
		base64.StdEncoding,
		base64.RawStdEncoding,
		base64.URLEncoding,
		base64.RawURLEncoding,
	}

	var err error
	for _, e := range encodings {
		var decoded []byte
		decoded, err = e.DecodeString(s)
		if err == nil {
			return decoded, nil
		}
	}
	
	return nil, err
}