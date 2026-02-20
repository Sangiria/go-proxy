package links

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func FetchVLESSLinks(url string) ([]string, error){
	response, err := http.Get(url)
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