package player

import (
	"fmt"
	"net/http"
	"time"
)

const endpoint = "https://api.spotify.com/v1/me/player/pause"

func Pause(client *http.Client, d time.Duration) error {
	req, err := http.NewRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		return err
	}

	timer := time.NewTimer(d)
	<-timer.C

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Status Code Error: response status code is %d", resp.StatusCode)
	}

	return nil
}
