package file

import (
	"core/models"
	"encoding/json"
	"os"
)

func SaveState(state *models.State) error {
	file, err := os.OpenFile("./state/state.json", os.O_WRONLY|os.O_TRUNC, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")

	if err := enc.Encode(state); err != nil {
		return err
	}

	return nil
}

//HandleAdd saves nodes and subscription if exists to file, sends [ok] if successfull (for now)
func LoadState() (*models.State, error) {
	//read state file if not create file

	if _, err := os.Stat("./state"); os.IsNotExist(err) {
		err = os.Mkdir("state", 0755)
		if err != nil {
			return nil, err
		}
	}

	if _, err := os.Stat("./state/state.json"); os.IsNotExist(err) {
		//create new file
		new_state := models.State{
			ActiveNodeId: "",
			Subscriptions: make(map[string]*models.Subscription),
			Nodes: make(map[string]*models.Node),
		}

		data, _ := json.MarshalIndent(new_state, "", "\t")
		if err = os.WriteFile("./state/state.json", data, 0600); err != nil {
			return nil, err
		}

		return &new_state, nil
	} else if err != nil {
		return nil, err
	}

	var state models.State

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