package main

import (
	"log"
	"os"
)

func main() {
	stopTime, err := AtoTime(os.Args[1])
	if err != nil || stopTime == -1 {
		log.Println(err, stopTime)
		os.Exit(1)
	}
	conf := NewConfig()
	result := NewClinet(conf)
	err = Pause(result.Client, stopTime)
	if err != nil {
		log.Println("Failed to pause", err)
	}
}
