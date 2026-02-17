package links

import (
	"core/models"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type parseResult struct {
	URLs			[]*url.URL
	SourseType		models.SourceType
}

//parsing the links the result is a ParseResult struct
func ParseURL(s_url string) (*parseResult, error) {
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

func ParseVLESSLink(url *url.URL, url_q url.Values) (*models.Parsed, error) {
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