package handlers

import (
	"core/manager"
	"core/models"
	"core/service"
	"core/utils"
	"fmt"
	"strings"
)

func AddNodesHandler(url string) (string, error) {
	url_string := strings.TrimSpace(url)
	if url_string == "" {
		return "error", fmt.Errorf("empty url")
	}
	
	//get state variable
	state, err := service.LoadState()
	if err != nil {
		return "error", fmt.Errorf("something went wrong while loading state: %w", err)
	}
	
	//parseInput function call
	result, err := utils.ParseInput(url_string)
	
	n := len(result.URLs)
	//create nodes from parseResult urls
	if n > 1 {

	} else if n == 1 {
		new_node, err := manager.CreateNode(result.URLs[0], models.Source{
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
	if err = service.SaveState(state); err != nil {
		return "error", err
	}

	return "ok", nil
}