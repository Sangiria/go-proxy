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
	URLs			[]*url.URL
	SourseType		models.SourceType
	Subscription	models.Subscription
}

func SaveState(state *models.State) error {
	data, _ := json.MarshalIndent(state, "", "\t")
	if err := os.WriteFile("./state/state.json", data, 0600); err != nil {
		return err
	}

	return nil
}

//HandleAdd saves nodes and subscription if exists to file, sends [ok] if successfull (for now)
func LoadState() (*models.State, error) {
	//read state file if not create file

	if _, err := os.Stat("./state"); os.IsNotExist(err) {
		err = os.Mkdir("state", 0755)
		if err != nil {
			return nil, err
		}
	}

	if _, err := os.Stat("./state/state.json"); os.IsNotExist(err) {
		//create new file
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

	//read existing file
	file, err := os.ReadFile("./state/state.json")
	if err != nil {
		return nil, err
	}

	//save to variable (struct State)
	err = json.Unmarshal(file, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func HandleAdd(url_string string) (string, error) {
	if url_string == "" {
		return "error", fmt.Errorf("empty url")
	}
	
	//get state variable
	state, err := LoadState()
	if err != nil {
		return "error", fmt.Errorf("something went wrong while loading state: %w", err)
	}
	
	//parseInput function call
	result, err := ParseInput(url_string)
	
	n := len(result.URLs)
	//create nodes from parseResult urls
	if n > 1 {

	} else if n == 1 {
		new_node, err := CreateNode(result.URLs[0], models.Source{
			Type: result.SourseType,
		})

		if err != nil {
			return "error", err
		}

		//update struct State
		state.Nodes = append(state.Nodes, new_node)
	} else {
		return "error", fmt.Errorf("nothing to add")
	}

	//update file
	if err = SaveState(state); err != nil {
		return "error", err
	}

	return "ok", nil
}

func CreateNode(u *url.URL, source models.Source) (*models.Node, error) {
	q_u := u.Query()
	parsed, err := parseVLESS(u, q_u)
	if err != nil {
		return nil, err
	}

	name := u.Fragment 
	if name == "" {
		name = u.Host
	}

	return &models.Node{
		ID: uuid.NewString(),
		Name: name,
		Source: source,
		URL: u.String(),
		Parsed: *parsed,
	}, nil
}

//parsing the links the result is a ParseResult struct
func ParseInput(s_url string) (*parseResult, error) {
	u, err := url.Parse(s_url)
	if err != nil {
		return nil, err
	}

	//reading url scheme
	switch u.Scheme {
	case "https":
		//fetch vless urls
		//return parseresult
	case "vless":
		//return parseResult
		return &parseResult{
			URLs: []*url.URL{u},
			SourseType: models.SourceManual,
		}, nil
	}

	return nil, fmt.Errorf("unsupported scheme %s", u.Scheme)
}

func FetchVLESSLinks(u *url.URL) {

}

func CreateSubscription() {

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