# spotify-timer
Dead simply sleep timer for spotify

# What is this 
As this title, this app is alarm for spotify

# Why 
I just wanted to controll spotify playback from terminal

# Setup 
At first, please go to [here](https://developer.spotify.com/dashboard/) to create Client ID and Client Secret 
After that, please set the values into ``` .env ``` file

```
go get github.com/justym/spotify-timer
cd $GOPATH/src/github.com/justym/spotify-timer
go build
```

# Usage
```
./spotify-timer [minutes]
```




