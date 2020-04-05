package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/justym/spotify-timer/cmd/spotify-timer/server"
	o "github.com/justym/spotify-timer/lib/oauth2"
	"github.com/justym/spotify-timer/lib/player"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8080", "Default is :8080")
}

func main() {

	flag.Parse()

	http.HandleFunc("/", server.Home)
	http.HandleFunc("/callback", server.Authorize)

	var client *http.Client
	go func() {
		url := o.GetRedirectURL()
		fmt.Println("Please Log In to Spotify from this URL: ", url)
		token := o.AccessToken
		client = <-o.Ch
		err := player.Pause(token, client)
		if err != nil {
			log.Fatal(err)
		}
	}()

	http.ListenAndServe(port, nil)
}
