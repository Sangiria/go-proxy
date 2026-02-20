package handlers

import (
	"core/file"
	"core/links"
	"core/models"
	"fmt"
	"net/url"
	"strings"
)

func GetNodesHandler() {

}

func AddNodesHandler(s string) (string, error) {
	url_string := strings.TrimSpace(s)
	if url_string == "" {
		return "error", fmt.Errorf("empty url")
	}
	
	//get state variable
	state, err := file.LoadState()
	if err != nil {
		return "error", fmt.Errorf("something went wrong while loading state: %w", err)
	}
	
	//parseInput function call
	result, err := links.ParseURL(url_string)

	switch(result.SourseType) {
	case models.SourceManual:
		if err := state.AddNodeFromURL(result.URLs[0], &models.Source{
			Type: result.SourseType,
		}); err != nil {
			return "error", err
		}

	case models.SourceSubscription:
		//create sub id
		sub_key := links.GenerateID(url_string)
		
		//check if sub exist
		_, found := state.Subscriptions[sub_key]
		if found {
			return "error", fmt.Errorf("subscription already exists")
		}

		//if not create
		name, _ := url.Parse(url_string)
		state.Subscriptions[sub_key] = &models.Subscription{
			Name: name.Host,
			URL: url_string,
		}

		for _, url := range result.URLs {
			if err := state.AddNodeFromURL(url, &models.Source{
				Type: models.SourceSubscription,
				SubscriptionID: sub_key,
			}); err != nil {
				continue
			}
		}
	}

	//update file
	if err = file.SaveState(state); err != nil {
		return "error", err
	}

	return "ok", nil
}