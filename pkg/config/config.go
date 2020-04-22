package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var config *oauth2.Config

const scopeUMPS = "user-modify-playback-state"

func NewConfig() *oauth2.Config {
	LoadEnv()
	id, key := GetValues()
	if config == nil {
		config = &oauth2.Config{
			ClientID:     id,
			ClientSecret: key,
			Endpoint:     spotify.Endpoint,
			Scopes:       []string{scopeUMPS},
		}
	}

	return config
}
