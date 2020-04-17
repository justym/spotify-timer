package main

import (
	"log"
	"os"

	a2ns "github.com/nmrshll/oauth2-noserver"
	"golang.org/x/oauth2"
)

var authorizedClient *a2ns.AuthorizedClient

func NewClinet(conf *oauth2.Config) *a2ns.AuthorizedClient {
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
