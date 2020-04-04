package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/justym/spotify-timer/controller"
	o "github.com/justym/spotify-timer/oauth2"
	"github.com/justym/spotify-timer/player"
)

func main() {
	var port string
	flag.StringVar(&port, "port", ":8080", "Default is :8080")
	flag.Parse()

	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/callback", controller.Authorize)

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
