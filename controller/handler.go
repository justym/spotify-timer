package controller

import (
	"fmt"
	"log"
	"net/http"
)

//Authorize is http handler for Spotify Authentication
func Authorize(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authenticate Handler")
	log.Println("Got Request for /auth")
}

//Home is http handler for Home
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Handler")
	log.Println("Got Request for /")
}
