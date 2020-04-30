# spotify-timer
Dead simple sleep timer for spotify

# Why 
I just wanted to use sleep timer from cli

# Setup 
1. Please go to [here](https://developer.spotify.com/dashboard/) to create Client ID and Client Secret 
2. Please set redirect uri to http://127.0.0.1:14565/oauth/callback in your APP
3. Please set these values like below.
``` sh
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
``` sh
./spotify-timer -d=[duration]
## NOTE: duration format is following.
##   3s ... 3 seconds
##   10m ... 10 minutes
```

# Welcome
- Pull Requests
- Issues
- Code Reviews

