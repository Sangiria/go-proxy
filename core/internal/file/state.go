package file

import (
	"core/models"
	"encoding/json"
	"os"
)

type State struct {
	ActiveNodeId			string							`json:"active_node"`
	Subscriptions			map[string]*models.Subscription	`json:"subscriptions"`
	Manual					map[string]*models.Node			`json:"manual"`
	ItemsOrder				[]string						`json:"items_order"`
}

func SaveState(s *State) error {
	tmp_path := "./state/state.json.tmp"
    final_path := "./state/state.json"

    f, err := os.Create(tmp_path)
    if err != nil { return err }
    defer f.Close()

    enc := json.NewEncoder(f)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")

    if err := enc.Encode(s); err != nil { 
		return err 
	}
    f.Close()

    return os.Rename(tmp_path, final_path)
}

//HandleAdd saves nodes and subscription if exists to file
func LoadState() (*State, error) {
	//read state file if not create file

	if _, err := os.Stat("./state"); os.IsNotExist(err) {
		err = os.Mkdir("state", 0755)
		if err != nil {
			return nil, err
		}
	}

	if _, err := os.Stat("./state/state.json"); os.IsNotExist(err) {
		//create new file
		new_state := State{
			ActiveNodeId: "",
			Subscriptions: make(map[string]*models.Subscription),
			Manual: make(map[string]*models.Node),
			ItemsOrder: make([]string, 0),
		}

		data, _ := json.MarshalIndent(new_state, "", "\t")
		if err = os.WriteFile("./state/state.json", data, 0600); err != nil {
			return nil, err
		}

		return &new_state, nil
	} else if err != nil {
		return nil, err
	}

	var state State

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