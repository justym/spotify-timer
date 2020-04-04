package main

import (
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
}


