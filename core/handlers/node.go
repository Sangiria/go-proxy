package handlers

import (
	"core/file"
	"core/links"
	"core/models"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
	t0 := time.Now()
	if err = file.SaveState(state); err != nil {
		return "error", err
	}
	dt := time.Since(t0)
	info, err := os.Stat("./state/state.json")
	if err != nil {
		log.Fatal(err)
	}

	size := info.Size()

	fmt.Printf("Size: %d bytes (%.2f KiB)\n", size, float64(size)/1024)
	fmt.Printf("Time: %v\n", dt)

	speed := float64(size) / dt.Seconds() / 1024
	fmt.Printf("Speed: %.2f KiB/s\n", speed)

	return "ok", nil
}