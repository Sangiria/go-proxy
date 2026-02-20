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

func AddNodesHandler(url string) (string, error) {
	url_string := strings.TrimSpace(url)
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
		if err := state.AddNodeFromURL(result.URLs[0], result.SourseType); err != nil {
			return "error", err
		}

	case models.SourceSubscription:
		
		// for _, url := range result.URLs {
		// 	if err := state.AddNodeFromURL(url, result.SourseType); err != nil {
		// 		continue
		// 	}
		// }
	}

	//update file
	if err = file.SaveState(state); err != nil {
		return "error", err
	}

	return "ok", nil
}