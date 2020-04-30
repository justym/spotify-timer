package auth

import (
	"log"

	"github.com/justym/spotify-timer/pkg/config"
	a2ns "github.com/nmrshll/oauth2-noserver"
)

func NewClinet() *a2ns.AuthorizedClient {
	conf := config.NewConfig()
	if conf == nil {
		log.Fatal("config is nil")
	}

	authorizedClient, err := a2ns.AuthenticateUser(conf)
	if err != nil {
		log.Fatal(err)
	}

	return authorizedClient
}
