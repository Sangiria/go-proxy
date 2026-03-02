package links

import (
	"core/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func ParseURLToNode(s string, source *models.Source) (*models.Node, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	//check for valid node url
	parsed, err := ParseNodeURL(u)
	if err != nil {
		return nil, err
	}

	name := u.Fragment
	if name == "" {
		name = u.Host
	}

	return &models.Node{
		Name: name,
		Source: *source,
		Parsed: *parsed,
	}, nil
}

func ParseNodeURL(u *url.URL) (*models.Parsed, error) {
	var (
		extra json.RawMessage
		u_q = u.Query()
	)

	uuid_str := strings.TrimSpace(u.User.Username())
	if uuid_str == "" {
		return nil, errors.New("uuid required")
	}
	if err := uuid.Validate(uuid_str); err != nil {
		return nil, errors.New("invalid uuid")
	}

	host, str_port := u.Hostname(), u.Port()
	if host == "" || str_port == "" {
    	return nil, errors.New("missing host or port")
  	}

	port, err := strconv.Atoi(str_port)
  	if err != nil || port < 1 || port > 65535 {
		return nil, errors.New("invalid port")
	}

	raw_extra := u_q.Get("extra")
	if raw_extra != "" {
		extra, err = parseExtra(raw_extra)
		if err != nil {
			return nil, err
		}
	}

	return &models.Parsed{
		Type: u.Scheme,
		Address: host,
		Port: uint16(port),
		UUID: uuid_str,
		Transport: u_q.Get("type"),
		Security: u_q.Get("security"),
		Sni: u_q.Get("sni"),
		Fp: u_q.Get("fp"),
		Pbk: u_q.Get("pbk"),
		Sid: u_q.Get("sid"),
		Flow: u_q.Get("flow"),
		Host: u_q.Get("host"),
		Path: u_q.Get("path"),
		XHTTPMode: u_q.Get("mode"),
		XHTTPExtra: extra,
	}, nil
}

func parseExtra(r string) (json.RawMessage, error) {
	decoded, err := url.QueryUnescape(r)
	if err != nil {
		return nil, fmt.Errorf("extra url decode failed: %s", err)
	}

	if !json.Valid([]byte(decoded)) {
		return nil, fmt.Errorf("JSON is not valid")
	}

	return json.RawMessage(decoded), nil
}