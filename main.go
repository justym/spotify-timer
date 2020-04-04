package main

import (
<<<<<<< HEAD
	"flag"
	"net/http"

	"github.com/justym/spotify-timer/controller"
)

func main() {
	var port string
	flag.StringVar(&port, "port", ":8080", "Default is :8080")
	flag.Parse()

	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/auth", controller.Authorize)
	http.ListenAndServe(port, nil)
=======
	own "github.com/justym/spotify-timer/oauth2"
	"fmt"
)

func main(){
	fmt.Println(own.Config.RedirectURL)
>>>>>>> 1c79828c5b5d8496d164401906da99c4f9e380b5
}
