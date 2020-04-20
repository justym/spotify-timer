# spotify-timer
Dead simple sleep timer for spotify

# Why 
I just wanted to use sleep timer from cli

# Setup 
1. Please go to [here](https://developer.spotify.com/dashboard/) to create Client ID and Client Secret 
2. Please set redirect uri to http://127.0.0.1:14565/oauth/callback
2. Please set the values into ``` .env ``` file
3. Please run the commands below
```
go get github.com/justym/spotify-timer
cd $GOPATH/src/github.com/justym/spotify-timer
go build
```

# Usage
```
./spotify-timer [minutes]
```

# Finally
Pull Request welcome :dizzy:




