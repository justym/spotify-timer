# spotify-timer
Dead simple sleep timer for spotify

# Why 
I just want to use sleep timer from cli

# Setup 
1. Please go to [here](https://developer.spotify.com/dashboard/) to create Client ID and Client Secret 
2. Please set redirect uri to http://127.0.0.1:14565/oauth/callback in your APP
3. Please set these values like below.
```
export SPOTIFY_TIMER_ID=your_client_id
export SPOTIFY_TIMER_KEY=your_client_key
```
4. Please run the commands below
``` sh
go get github.com/justym/spotify-timer
cd $GOPATH/src/github.com/justym/spotify-timer
go build
```

# Usage
```sh
./spotify-timer -h #help

./spotify-timer [minutes]
./spotify-timer -m=1m (a minute)
./spotify-timer (default 10m)
```

# Built with
- github.com/spf13/cobra
- github.com/nmrshll/oauth2-noserver
- github.com/fatih/color

# Welcome
- Pull Requests
- Issues
- Code Reviews





