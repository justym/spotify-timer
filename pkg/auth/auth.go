package auth

import (
	a2ns "github.com/nmrshll/oauth2-noserver"

	"github.com/justym/spotify-timer/pkg/config"
)

func NewClient() (*a2ns.AuthorizedClient, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	return a2ns.AuthenticateUser(conf)
}
