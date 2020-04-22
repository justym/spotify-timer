package config

import (
	"os"
)

var (
	clientID     string
	clientSecret string
)

//LoadEnv set values from .env file
func LoadEnv() {
	clientID = os.Getenv("SPOTIFY_TIMER_ID")
	clientSecret = os.Getenv("SPOTIFY_TIMER_KEY")
}

//GetValues returns client id, client secret key
func GetValues() (string, string) {
	return clientID, clientSecret
}
