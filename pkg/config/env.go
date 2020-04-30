package config

import (
	"fmt"
	"os"
)

var (
	clientID     string
	clientSecret string
)

func init() {
	clientID = os.Getenv("SPOTIFY_TIMER_ID")
	clientSecret = os.Getenv("SPOTIFY_TIMER_KEY")
}

//GetValues returns client id, client secret key
func GetValues() (string, string, error) {
	if clientID == "" || clientSecret == "" {
		err := fmt.Errorf("SPOTIFY_TIMER_ID or SPOTIFY_TIMER_KEY is empty")
		return clientID, clientSecret, err
	}
	return clientID, clientSecret, nil
}
