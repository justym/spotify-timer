package auth

import (
	"log"
	"os"

	"github.com/justym/spotify-timer/pkg/config"
	a2ns "github.com/nmrshll/oauth2-noserver"
)

var authorizedClient *a2ns.AuthorizedClient

func NewClinet() *a2ns.AuthorizedClient {
	conf := config.NewConfig()
	if conf == nil {
		log.Fatal("config is nil")
	}
	if authorizedClient == nil {
		authorizedClient, err := a2ns.AuthenticateUser(conf)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		return authorizedClient
	}

	return authorizedClient
}
