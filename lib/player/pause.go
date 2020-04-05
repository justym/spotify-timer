package player

import (
	"errors"
	"log"
	"net/http"
)

const pauseURL = "https://api.spotify.com/v1/me/player/pause"

//Pause pauses playback
func Pause(accessToken string, client *http.Client) error {
	req, err := http.NewRequest(http.MethodPut, pauseURL, nil)
	if err != nil {
		log.Println(err)
		return errors.New("Failed to create New Request in PauseOpt")
	}
	req.Header.Set("Authorization", "Bearer "+ accessToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return errors.New("Falied to do clinet.Do(req) in PauseOpt")
	}
	if resp.StatusCode != 204 {
		log.Println("Invalid Status Code: %d", resp.StatusCode)
		return errors.New("Invalid Status Code: " + string(resp.StatusCode))
	}

	return nil
}
