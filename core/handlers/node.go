package handlers

import (
	"core/file"
	"core/links"
	"core/models"
	"fmt"
	"strings"
)

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
		new_node, err := links.CreateNode(result.URLs[0], models.Source{
			Type: result.SourseType,
		})

		if err != nil {
			return "error", err
		}

		//update struct State
		node_key := links.GenerateDeterministicID(new_node.Parsed)
		_, found := state.Nodes[node_key]
		if found {
			return "error", fmt.Errorf("node already exist")
		}

		state.Nodes[node_key] = new_node

	case models.SourceSubscription:
		//create subscription
	}

	//update file
	if err = file.SaveState(state); err != nil {
		return "error", err
	}

	return "ok", nil
}