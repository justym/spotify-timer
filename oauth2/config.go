package oauth2

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

//Config is configuration for Oauth2
var (
	Config = oauth2.Config{
		ClientID:     spotifyID,
		ClientSecret: spotifyKey,
		Endpoint:     spotify.Endpoint,
		RedirectURL:  redirectURI,
		Scopes: []string{
			"user-modify-playback-state",
		},
	}

	client = make(chan *http.Client)
)

const (
	spotifyID   = ""
	spotifyKey  = ""
	redirectURI = ""
)

//GetToken get token from spotify
func GetToken(code string) (*oauth2.Token, error) {
	return Config.Exchange(oauth2.NoContext, code)
}

//GetRedirectURL return Auth Code URL
func GetRedirectURL() string {
	return Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
}
