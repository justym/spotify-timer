package main

import (
	own "github.com/justym/spotify-timer/oauth2"
	"fmt"
)

func main(){
	fmt.Println(own.Config.RedirectURL)
}
