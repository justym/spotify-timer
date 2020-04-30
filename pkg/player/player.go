package player

import (
	"fmt"
	"net/http"
	"time"
)

const endpoint = "https://api.spotify.com/v1/me/player/pause"

func Pause(client *http.Client, d time.Duration) error {
	timer := time.NewTimer(d)
	<-timer.C

	req, err := http.NewRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Status Code Error: %d", resp.StatusCode)
	}

	return nil
}
