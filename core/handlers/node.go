package handlers

import (
	"core/file"
	"core/links"
	"core/models"
	"fmt"
	"strings"
)

func GetNodesHandler() {

}

func AddNodesHandler(s string) (string, error) {
	//remove in the future
	url_string := strings.TrimSpace(s)
	if url_string == "" {
		return "error", fmt.Errorf("empty url")
	}
	
	//get state variable
	state, err := file.LoadState()
	if err != nil {
		return "error", fmt.Errorf("something went wrong while loading state: %w", err)
	}
	
	//parcing url and if https fetching nodes
	result, err := links.ParseURL(url_string)

	switch(result.SourseType) {
	case models.SourceManual:
		if err := state.AddNodeFromURL(result.URLs[0], &models.Source{
			Type: result.SourseType,
		}); err != nil {
			return "error", err
		}

	case models.SourceSubscription:
		sub_key, err := state.AddSubscriptionFromURL(url_string)
		if err != nil {
			return "error", err
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