package service

import (
	"core/api"
	"core/internal/links"
	"core/models"
	"errors"

	"github.com/google/uuid"
)

func UpdateNodeFromForm(node *models.Node, node_form *api.NodeForm) error {
	if node_form.Name != nil {
		node.Name = *node_form.Name
	}

	if node_form.Address != nil {
		node.Parsed.Address = *node_form.Address
	}

	if node_form.Port != nil {
		node.Parsed.Port = uint16(*node_form.Port)
	}

	if node_form.Uuid != nil {
		if err := uuid.Validate(*node_form.Uuid); err != nil {
        	return errors.New("invalid uuid")
    	}
		node.Parsed.UUID = *node_form.Uuid
	}

	if node_form.Transport != nil {
		node.Parsed.Transport = *node_form.Transport
	}

	if node_form.Security != nil {
		node.Parsed.Security = *node_form.Security
	}

	if node_form.Sni != nil {
		node.Parsed.Sni = *node_form.Sni
	}

	if node_form.Fp != nil {
		node.Parsed.Fp = *node_form.Fp
	}

	if node_form.Pbk != nil {
		node.Parsed.Pbk = *node_form.Pbk
	}

	if node_form.Sid != nil {
		node.Parsed.Sid = *node_form.Sid
	}

	if node_form.Mode != nil {
		node.Parsed.XHTTPMode = *node_form.Mode
	}

	if node_form.Extra != nil {
		if *node_form.Extra != "" {
			extra, err := links.ParseExtra(*node_form.Extra)
			if err != nil {
				return err
			}
			node.Parsed.XHTTPExtra = extra
		} else {
			node.Parsed.XHTTPExtra = nil
		}
	}
	
	return nil
}

func UpdateSubscriptionFromForm(sub *models.Subscription, sub_form *api.SubscriptionForm) {
	if sub_form.Name != nil {
		sub.Name = *sub_form.Name
	}

	if sub_form.Url != nil {
		sub.URL = *sub_form.Url
	}
}
