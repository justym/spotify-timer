package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	o "github.com/justym/spotify-timer/oauth2"
)

//Authorize is http handler for Spotify Authentication
func Authorize(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := o.GetToken(code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		os.Exit(1)
	}
	o.AccessToken = token.AccessToken
	o.Ch <- o.Config.Client(context.Background(), token)

	fmt.Fprintf(w, "Authenticate Handler")
	log.Println("Got Request for /auth")
}

//Home is http handler for Home
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home Handler")
	log.Println("Got Request for /")
}
