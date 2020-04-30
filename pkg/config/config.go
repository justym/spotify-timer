package config

import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

const scopeUMPS = "user-modify-playback-state"

func NewConfig() *oauth2.Config {
	id, key, err := GetValues()
	if err != nil {
		log.Fatal(err)
	}

	config := &oauth2.Config{
		ClientID:     id,
		ClientSecret: key,
		Endpoint:     spotify.Endpoint,
		Scopes:       []string{scopeUMPS},
	}

	return config
}
