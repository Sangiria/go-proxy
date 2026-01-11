package manager

import (
	"core/models"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type parseResult struct {
	URLs			[]string
	SourseType		models.SourceType
	Subscription	models.Subscription
}

//HandleAdd saves nodes and subscription if exists to file, sends [ok] if successfull
func LoadState() (*models.State, error) {
	//read state file if not create file

	if _, err := os.Stat("./state"); os.IsNotExist(err) {
		err = os.Mkdir("state", 0755)
		if err != nil {
			return nil, err
		}
	}

	if _, err := os.Stat("./state/state.json"); os.IsNotExist(err) {
		new_state := models.State{
			ActiveNodeId: "",
			Subscriptions: []*models.Subscription{},
			Nodes: []*models.Node{},
		}

		data, _ := json.MarshalIndent(new_state, "", "\t")
		if err = os.WriteFile("./state/state.json", data, 0600); err != nil {
			return nil, err
		}

		return &new_state, nil
	} else if err != nil {
		return nil, err
	}

	var state models.State

	file, err := os.ReadFile("./state/state.json")
	if err != nil {
		return nil, err
	}

	//save to variable (struct State)
	err = json.Unmarshal(file, state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func HandleAdd(url_string string) (string, error) {
	if url_string == "" {
		return "error", fmt.Errorf("empty url")
	}
	
	_, err := LoadState()
	if err != nil {
		return "error", fmt.Errorf("something went wrong while loading state: %w", err)
	}
	
	//parseInput function call
	//get parseResult

	//create nodes from parseResult urls
	//update struct State
	//update file

	return "ok", nil
}

func CreateSubscription() {

}

func CreateNode(url string, source models.Source) {

}

func FetchVLESSLinks() {

}

func parseVLESS(url *url.URL, url_q url.Values) (*models.Parsed, error) {
	var (
		transport models.Transport
		security models.Security
		path string
		mode string
		extra json.RawMessage
	)

	uuid_str := strings.TrimSpace(url.User.Username())
	if uuid_str == "" {
		return nil, errors.New("uuid required")
	}
	if err := uuid.Validate(uuid_str); err != nil {
		return nil, errors.New("invalid uuid")
	}

	host, str_port := url.Hostname(), url.Port()
	if host == "" || str_port == "" {
    	return nil, errors.New("missing host or port")
  	}

	port, err := strconv.Atoi(str_port)
  	if err != nil || port < 1 || port > 65535 {
		return nil, errors.New("invalid port")
	}

	switch url_q.Get("type") {
	case "tcp", "":
		transport = models.TransportTCP
	case "xhttp":
		transport = models.TransportXHTTP
		p := url_q.Get("path")
		if p == "" {
			path = "/"
		} else if !strings.HasPrefix(p, "/"){
			path = "/" + p
		}

		m := url_q.Get("mode")
		if m == "" {
			mode = "auto"
		} else {
			switch m {
			case "auto", "packet", "stream":
				mode = m
			default:
				return nil, fmt.Errorf("unsupported xhttp mode %s", m)
			}
		}

		raw_extra := url_q.Get("extra")
		if raw_extra != "" {
			extra, err = parseExtra(raw_extra)
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("%s is unsupported transport", url_q.Get("type"))
	}

	switch url_q.Get("security") {
	case "reality":
		security = models.SecurityReality
		
		required := []string{"sni", "fp", "pbk"}
		for _, key := range required {
			if strings.TrimSpace(url_q.Get(key)) == "" {
				return nil, fmt.Errorf("missing %s", key)
    		}
		}

		if url_q.Get("sid") != "" {
			if _, err := hex.DecodeString(url_q.Get("sid")); err != nil {
				return nil, fmt.Errorf("invalid sid")
			}
		}
	case "":
		security = models.SecurityNone
	default:
		return nil, errors.New("unsupported security")
	}

	return &models.Parsed{
		Address: host,
		Port: uint16(port),
		UUID: uuid_str,
		Transport: transport,
		Security: security,
		Sni: url_q.Get("sni"),
		Fp: url_q.Get("fp"),
		Pbk: url_q.Get("pbk"),
		Sid: url_q.Get("sid"),
		Flow: url_q.Get("flow"),
		Host: url_q.Get("host"),
		Path: path,
		XHTTPMode: mode,
		XHTTPExtra: extra,
	}, nil
}

func parseExtra(raw string) (json.RawMessage, error) {
	decoded, err := url.QueryUnescape(raw)
	if err != nil {
		return nil, fmt.Errorf("extra url decode failed: %s", err)
	}

	if !json.Valid([]byte(decoded)) {
		return nil, fmt.Errorf("JSON is not valid")
	}

	return json.RawMessage(decoded), nil
}