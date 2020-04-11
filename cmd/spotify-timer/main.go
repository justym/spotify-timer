package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/justym/spotify-timer/cmd/spotify-timer/server"
	o "github.com/justym/spotify-timer/lib/oauth2"
	"github.com/justym/spotify-timer/lib/player"
)

var port string
var m int

func init() {
	flag.StringVar(&port, "port", ":8080", "[Default] :8080")
	flag.IntVar(&m, "m", 15, "[Default] 15minute")
}

func main() {
	flag.Parse()

	http.HandleFunc("/callback", server.Authorize)

	tstop := time.Duration(m) * time.Minute
	go pause(tstop)

	http.ListenAndServe(port, nil)
}

func pause(d time.Duration) {
	url := o.GetRedirectURL()
	fmt.Println("Please Log in to spotify from this URL: ", url)
	token := o.AccessToken
	client := <-o.Ch
	t := time.NewTimer(d)
	<-t.C
	err := player.Pause(token, client)
	if err != nil {
		log.Println(err)
	}
}
