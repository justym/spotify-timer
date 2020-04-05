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

	AccessToken = ""
	Ch          = make(chan *http.Client)
)

const (
	spotifyID   = "608c6b65acf844ec8951b15c2b3a5e47"
	spotifyKey  = "2b173db76387433b8e4b384e15050134"
	redirectURI = "http://localhost:8080/callback"
)

//GetToken get token from spotify
func GetToken(code string) (*oauth2.Token, error) {
	return Config.Exchange(oauth2.NoContext, code)
}

//GetRedirectURL return Auth Code URL
func GetRedirectURL() string {
	return Config.AuthCodeURL("state")
}
