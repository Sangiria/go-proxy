package manager

import (
	"core/models"
	"core/utils"
	"net/url"
)

func CreateNode(u *url.URL, source models.Source) (*models.Node, error) {
	q_u := u.Query()
	parsed, err := utils.ParseVLESS(u, q_u)
	if err != nil {
		return nil, err
	}

	name := u.Fragment 
	if name == "" {
		name = u.Host
	}

	return &models.Node{
		Name: name,
		Source: source,
		URL: u.String(),
		Parsed: *parsed,
	}, nil
}