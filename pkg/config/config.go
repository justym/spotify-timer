package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

const scopeUMPS = "user-modify-playback-state"

func NewConfig() (*oauth2.Config, error) {
	envs, err := loadEnv()
	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     envs.clientID,
		ClientSecret: envs.clientKey,
		Endpoint:     spotify.Endpoint,
		Scopes:       []string{scopeUMPS},
	}, nil
}
