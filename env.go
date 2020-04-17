package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	clientID     string
	clientSecret string
)

//LoadEnv set values from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	clientID = os.Getenv("ID")
	clientSecret = os.Getenv("KEY")
}

//GetValues returns client id, client secret key
func GetValues() (string, string) {
	return clientID, clientSecret
}
