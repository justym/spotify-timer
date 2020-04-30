package config

import (
	"errors"
	"os"
)

type envs struct {
	clientID  string
	clientKey string
}

// loadEnv set & return values from .env file
func loadEnv() (*envs, error) {
	clientID := os.Getenv("SPOTIFY_TIMER_ID")
	clientKey := os.Getenv("SPOTIFY_TIMER_KEY")

	if clientID == "" || clientKey == "" {
		return nil, errors.New("SPOTIFY_TIMER_ID or SPOTIFY_TIMER_KEY is empty")
	}

	return &envs{
		clientID:  clientID,
		clientKey: clientKey,
	}, nil
}
